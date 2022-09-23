package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/SKF/go-asset-component-client/client"
	"github.com/SKF/go-asset-component-client/models"
	"github.com/SKF/go-utility/v2/uuid"
)

var _ client.API = &mockClient{nil}

type mockClient struct {
	*mock.Mock
}

func NewMockClient() *mockClient {
	return &mockClient{
		&mock.Mock{},
	}
}

func (c *mockClient) GetAssetComponents(ctx context.Context, assetID uuid.UUID, componentType ...string) (models.Components, error) {
	args := c.Called(ctx, assetID, componentType)
	return args.Get(0).(models.Components), args.Error(1)
}
