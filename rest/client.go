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

	GetAllComponentRelations(context.Context, uuid.UUID, int, string) (models.GetComponentRelationsResponse, error)
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

func (c *client) GetAllComponentRelations(ctx context.Context, id uuid.UUID, limit int, continuationToken string) (models.GetComponentRelationsResponse, error) {
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
