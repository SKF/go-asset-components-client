package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/SKF/go-utility/v2/uuid"

	"github.com/SKF/go-asset-components-client/rest"
	"github.com/SKF/go-asset-components-client/rest/models"
)

var _ rest.Client = &AssetComponentsMock{}

type AssetComponentsMock struct {
	*mock.Mock
}

func NewAssetComponentsMock() *AssetComponentsMock {
	return &AssetComponentsMock{&mock.Mock{}}
}

func (c *AssetComponentsMock) GetComponentsByAsset(ctx context.Context, id uuid.UUID, filter rest.GetComponentsFilter) ([]models.Component, error) {
	args := c.Called(ctx, id, filter)
	return args.Get(0).([]models.Component), args.Error(1)
}

func (c *AssetComponentsMock) GetComponentRelations(ctx context.Context, id uuid.UUID, limit int, continuationToken string) (models.GetComponentRelationsResponse, error) {
	args := c.Called(ctx, id, limit, continuationToken)
	return args.Get(0).(models.GetComponentRelationsResponse), args.Error(1)
}

func (c *AssetComponentsMock) GetRelatedComponents(ctx context.Context, id uuid.UUID, limit int, source, relationType, continuationToken string) (models.GetRelatedComponentsResponse, error) {
	args := c.Called(ctx, id, limit, source, relationType, continuationToken)
	return args.Get(0).(models.GetRelatedComponentsResponse), args.Error(1)
}

func (c *AssetComponentsMock) CreateComponentRelation(ctx context.Context, id uuid.UUID, relation models.Relation) error {
	args := c.Called(ctx, id, relation)
	return args.Error(0)
}

func (c *AssetComponentsMock) DeleteComponentRelation(ctx context.Context, externalID, componentID uuid.UUID, source, relationType string) (err error) {
	args := c.Called(ctx, externalID, componentID, source, relationType)
	return args.Error(0)
}
