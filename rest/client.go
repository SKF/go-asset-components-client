package rest

import (
	"context"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"

	"github.com/SKF/go-asset-components-client/rest/models"
)

type GetComponentsFilter struct {
	Types []string
}

type Client interface {
	GetComponentsByAsset(context.Context, uuid.UUID, GetComponentsFilter) ([]models.Component, error)

	GetComponentRelations(context.Context, uuid.UUID, int, string) (models.GetComponentRelationsResponse, error)
	GetRelatedComponents(context.Context, uuid.UUID, int, string, string, string) (models.GetRelatedComponentsResponse, error)
	CreateComponentRelation(ctx context.Context, id uuid.UUID, relation models.Relation) error
	DeleteComponentRelation(ctx context.Context, externalID, componentID uuid.UUID, source, relationType string) (err error)
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

func (c *client) GetComponentRelations(ctx context.Context, id uuid.UUID, limit int, continuationToken string) (models.GetComponentRelationsResponse, error) {
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

func (c *client) GetRelatedComponents(ctx context.Context, id uuid.UUID, limit int, source, relationType, continuationToken string) (models.GetRelatedComponentsResponse, error) {
	request := rest.Get("/relations/{source}/{type}/{id}/components{?limit,continuation_token*}").
		Assign("source", source).
		Assign("type", relationType).
		Assign("id", id).
		Assign("limit", limit).
		Assign("continuation_token", continuationToken).
		SetHeader("Accept", "application/json")

	var response models.GetRelatedComponentsResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetRelatedComponentsResponse{}, err
	}

	return response, nil
}

func (c *client) CreateComponentRelation(ctx context.Context, id uuid.UUID, relation models.Relation) error {
	request := rest.Put("/components/{component}/relations").
		Assign("component", id).
		WithJSONPayload(relation).
		SetHeader("Accept", "application/json")

	if _, err := c.Do(ctx, request); err != nil {
		return err
	}

	return nil
}

func (c *client) DeleteComponentRelation(ctx context.Context, externalID, componentID uuid.UUID, source, relationType string) (err error) {
	request := rest.Delete("/relations/{source}/{type}/{id}/components/{component}").
		Assign("component", componentID).
		Assign("source", source).
		Assign("type", relationType).
		Assign("id", externalID).
		SetHeader("Accept", "application/json")

	if _, err = c.Do(ctx, request); err != nil {
		return err
	}

	return nil
}
