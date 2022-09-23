package client

import (
	"context"
	"fmt"
	"net/url"

	internalmodel "github.com/SKF/go-asset-component-client/internal/models"
	"github.com/SKF/go-asset-component-client/models"
	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"
)

var _ API = &client{nil}

type (
	API interface {
		GetAssetComponents(ctx context.Context, assetID uuid.UUID, componentType ...string) (models.Components, error)
	}
	client struct {
		*rest.Client
	}
)

func NewClient(opts ...rest.Option) *client {
	return &client{
		rest.NewClient(
			append([]rest.Option{
				WithStage(stages.StageProd),
			}, opts...)...,
		),
	}
}

func WithStage(stage string) rest.Option {
	if stage == stages.StageProd {
		return rest.WithBaseURL("https://api.asset-components.enlight.skf.com")
	}

	return rest.WithBaseURL(fmt.Sprintf("https://api.%s.asset-components.enlight.skf.com", stage))
}

func (c *client) GetAssetComponents(ctx context.Context, assetID uuid.UUID, componentType ...string) (models.Components, error) {
	params := url.Values{"type": componentType}
	req := rest.Get("/asset/{assetID}/components?{query}").
		Assign("assetID", assetID).
		Assign("query", params.Encode()).
		SetHeader("Accept", "application/json")

	var internalComponent internalmodel.GetAssetComponentsResponse
	if err := c.DoAndUnmarshal(ctx, req, &internalComponent); err != nil {
		return models.Components{}, err
	}

	res := models.Components{
		Components: make([]models.Component, 0, len(internalComponent.Components)),
	}

	if err := res.FromInternal(internalComponent); err != nil {
		return models.Components{}, err
	}

	return res, nil
}
