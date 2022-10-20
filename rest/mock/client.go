package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/SKF/go-asset-components-client/rest"
	"github.com/SKF/go-asset-components-client/rest/models"
	"github.com/SKF/go-utility/v2/uuid"
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

func (c *AssetComponentsMock) GetAllComponentRelations(ctx context.Context, id uuid.UUID, limit int, continuationToken string) (models.GetComponentRelationsResponse, error) {
	args := c.Called(ctx, id, limit, continuationToken)
	return args.Get(0).(models.GetComponentRelationsResponse), args.Error(1)
}
