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

func (c *AssetComponentsMock) GetComponent(ctx context.Context, id uuid.UUID) (models.Component, error) {
	args := c.Called(ctx, id)
	return args.Get(0).(models.Component), args.Error(1)
}

func (c *AssetComponentsMock) GetComponentsByAsset(ctx context.Context, id uuid.UUID, filter rest.GetComponentsFilter) ([]models.Component, error) {
	args := c.Called(ctx, id, filter)
	return args.Get(0).([]models.Component), args.Error(1)
}

func (c *AssetComponentsMock) CreateComponent(ctx context.Context, component models.Component) (models.Component, error) {
	args := c.Called(ctx, component)
	return args.Get(0).(models.Component), args.Error(1)
}

func (c *AssetComponentsMock) DeleteComponent(ctx context.Context, id uuid.UUID) error {
	args := c.Called(ctx, id)
	return args.Error(0)
}

func (c *AssetComponentsMock) UpdateComponent(ctx context.Context, component models.Component) (models.Component, error) {
	args := c.Called(ctx, component)
	return args.Get(0).(models.Component), args.Error(1)
}

func (c *AssetComponentsMock) GetComponentRelations(ctx context.Context, id uuid.UUID) ([]models.Relation, error) {
	args := c.Called(ctx, id)
	return args.Get(0).([]models.Relation), args.Error(1)
}

func (c *AssetComponentsMock) GetRelatedComponents(ctx context.Context, relation models.Relation) ([]models.RelatedComponent, error) {
	args := c.Called(ctx, relation)
	return args.Get(0).([]models.RelatedComponent), args.Error(1)
}

func (c *AssetComponentsMock) CreateComponentRelation(ctx context.Context, relation models.Relation, id uuid.UUID) error {
	args := c.Called(ctx, relation, id)
	return args.Error(0)
}

func (c *AssetComponentsMock) DeleteComponentRelation(ctx context.Context, relation models.Relation, id uuid.UUID) (err error) {
	args := c.Called(ctx, relation, id)
	return args.Error(0)
}

func (c *AssetComponentsMock) GetComponentSpeed(ctx context.Context, id uuid.UUID) (models.CalculatedSpeed, models.SpeedConfiguration, error) {
	args := c.Called(ctx, id)
	return args.Get(0).(models.CalculatedSpeed), args.Get(1).(models.SpeedConfiguration), args.Error(2) //nolint:gomnd
}

func (c *AssetComponentsMock) SetComponentSpeed(ctx context.Context, id uuid.UUID, speedConfiguration models.SpeedConfiguration) error {
	args := c.Called(ctx, id, speedConfiguration)
	return args.Error(0)
}

func (c *AssetComponentsMock) DeleteComponentSpeed(ctx context.Context, id uuid.UUID) error {
	args := c.Called(ctx, id)
	return args.Error(0)
}
