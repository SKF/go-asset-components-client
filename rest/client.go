package rest

import (
	"context"
	"fmt"
	"net/url"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"

	"github.com/SKF/go-asset-components-client/rest/models"
)

type GetComponentsFilter struct {
	Types []string
}

type Client interface {
	GetComponent(context.Context, uuid.UUID) (models.Component, error)
	GetComponentsByAsset(context.Context, uuid.UUID, GetComponentsFilter) ([]models.Component, error)
	CreateComponent(context.Context, models.Component) (models.Component, error)
	DeleteComponent(context.Context, uuid.UUID) error
	UpdateComponent(context.Context, models.Component) (models.Component, error)

	GetComponentRelations(context.Context, uuid.UUID) ([]models.Relation, error)
	GetRelatedComponents(context.Context, models.Relation) ([]models.RelatedComponent, error)
	CreateComponentRelation(context.Context, models.Relation, uuid.UUID) error
	DeleteComponentRelation(context.Context, models.Relation, uuid.UUID) (err error)

	GetComponentSpeed(context.Context, uuid.UUID) (models.GetComponentSpeedResponse, error)
	SetComponentSpeed(context.Context, uuid.UUID, models.SpeedConfiguration) (models.PutComponentSpeedResponse, error)
	DeleteComponentSpeed(context.Context, uuid.UUID) error
}

type client struct {
	*rest.Client
}

type Option = rest.Option

func NewClient(opts ...Option) Client {
	restClient := rest.NewClient(
		append([]Option{
			// Defaults to production stage if no option is supplied
			WithStage(stages.StageProd),
			rest.WithProblemDecoder(new(rest.BasicProblemDecoder)),
		}, opts...)...,
	)

	return &client{Client: restClient}
}

func (c *client) GetComponent(ctx context.Context, id uuid.UUID) (models.Component, error) {
	request := rest.Get("components/{component}").
		Assign("component", id).
		SetHeader("Accept", "application/json")

	var response models.Component
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Component{}, err
	}

	return response, nil
}

func (c *client) GetComponentsByAsset(ctx context.Context, id uuid.UUID, filter GetComponentsFilter) ([]models.Component, error) {
	request := rest.Get("assets/{asset}/components{?type*}").
		Assign("asset", id).
		Assign("type", filter.Types).
		SetHeader("Accept", "application/json")

	var response models.GetAssetComponentsResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return nil, err
	}

	return response.Components, nil
}

func (c *client) CreateComponent(ctx context.Context, component models.Component) (models.Component, error) {
	request := rest.Post("/components").
		WithJSONPayload(component).
		SetHeader("Accept", "application/json")

	var response models.Component
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Component{}, err
	}

	return response, nil
}

func (c *client) DeleteComponent(ctx context.Context, id uuid.UUID) error {
	request := rest.Delete("components/{component}").
		Assign("component", id).
		SetHeader("Accept", "application/json")

	_, err := c.Do(ctx, request)

	return err
}

func (c *client) UpdateComponent(ctx context.Context, component models.Component) (models.Component, error) {
	request := rest.Patch("components/{component}").
		Assign("component", *component.Id).
		WithJSONPayload(component).
		SetHeader("Accept", "application/json")

	var response models.Component
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Component{}, err
	}

	return response, nil
}

func (c *client) GetComponentRelations(ctx context.Context, id uuid.UUID) ([]models.Relation, error) {
	response, err := c.getComponentRelationsPage(ctx, id, 0, "")
	if err != nil {
		return nil, err
	}

	relations := response.Relations

	for response.Links.Next != nil {
		continuationToken, err := getContinuationToken(*response.Links.Next)
		if err != nil {
			return nil, err
		}

		response, err = c.getComponentRelationsPage(ctx, id, 0, continuationToken)
		if err != nil {
			return nil, err
		}

		relations = append(relations, response.Relations...)
	}

	return relations, nil
}

func (c *client) getComponentRelationsPage(ctx context.Context, id uuid.UUID, limit int, continuationToken string) (models.GetComponentRelationsResponse, error) {
	request := rest.Get("/components/{component}/relations{?limit,continuation_token*}").
		Assign("component", id).
		Assign("limit", limit).
		Assign("continuation_token", continuationToken).
		SetHeader("Accept", "application/json")

	var response models.GetComponentRelationsResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetComponentRelationsResponse{}, err
	}

	return response, nil
}

func (c *client) GetRelatedComponents(ctx context.Context, relation models.Relation) ([]models.RelatedComponent, error) {
	response, err := c.getRelatedComponentsPage(ctx, relation, 0, "")
	if err != nil {
		return nil, err
	}

	relatedComponents := response.Components

	for response.Links.Next != nil {
		continuationToken, err := getContinuationToken(*response.Links.Next)
		if err != nil {
			return nil, err
		}

		response, err = c.getRelatedComponentsPage(ctx, relation, 0, continuationToken)
		if err != nil {
			return nil, err
		}

		relatedComponents = append(relatedComponents, response.Components...)
	}

	return relatedComponents, nil
}

func (c *client) getRelatedComponentsPage(ctx context.Context, relation models.Relation, limit int, continuationToken string) (models.GetRelatedComponentsResponse, error) {
	request := rest.Get("/relations/{source}/{type}/{id}/components{?limit,continuation_token*}").
		Assign("source", relation.Source).
		Assign("type", relation.Type).
		Assign("id", relation.Id).
		Assign("limit", limit).
		Assign("continuation_token", continuationToken).
		SetHeader("Accept", "application/json")

	var response models.GetRelatedComponentsResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetRelatedComponentsResponse{}, err
	}

	return response, nil
}

func (c *client) CreateComponentRelation(ctx context.Context, relation models.Relation, id uuid.UUID) error {
	request := rest.Put("/components/{component}/relations").
		Assign("component", id).
		WithJSONPayload(relation).
		SetHeader("Accept", "application/json")

	_, err := c.Do(ctx, request)

	return err
}

func (c *client) DeleteComponentRelation(ctx context.Context, relation models.Relation, id uuid.UUID) error {
	request := rest.Delete("/relations/{source}/{type}/{id}/components/{component}").
		Assign("source", relation.Source).
		Assign("type", relation.Type).
		Assign("id", relation.Id).
		Assign("component", id).
		SetHeader("Accept", "application/json")

	_, err := c.Do(ctx, request)

	return err
}

func getContinuationToken(nextLink string) (string, error) {
	queryValues, err := url.Parse(nextLink)
	if err != nil {
		err = fmt.Errorf("unable to parse next Link:'%s' as an URL %w", nextLink, err)
		return "", err
	}

	continuationToken := queryValues.Query().Get("continuation_token")
	if continuationToken == "" {
		err = fmt.Errorf("expected query value continuation_token not found in next Link: %s", nextLink)
		return "", err
	}

	return continuationToken, nil
}

func (c *client) GetComponentSpeed(ctx context.Context, id uuid.UUID) (models.GetComponentSpeedResponse, error) {
	request := rest.Get("/components/{component}/speed").
		Assign("component", id).
		SetHeader("Accept", "application/json")

	var response models.GetComponentSpeedResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetComponentSpeedResponse{}, err
	}

	return response, nil
}

func (c *client) SetComponentSpeed(ctx context.Context, id uuid.UUID, speedConfiguraion models.SpeedConfiguration) (models.PutComponentSpeedResponse, error) {
	request := rest.Put("/components/{component}/speed").
		Assign("component", id).
		WithJSONPayload(models.PutComponentSpeedRequest{Configuration: speedConfiguraion}).
		SetHeader("Accept", "application/json")

	var response models.PutComponentSpeedResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.PutComponentSpeedResponse{}, err
	}

	return response, nil
}

func (c *client) DeleteComponentSpeed(ctx context.Context, id uuid.UUID) error {
	request := rest.Delete("/components/{component}/speed").
		Assign("component", id).
		SetHeader("Accept", "application/json")

	_, err := c.Do(ctx, request)

	return err
}
