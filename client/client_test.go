package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internalmodel "github.com/SKF/go-asset-component-client/internal/models"

	"github.com/SKF/go-asset-component-client/models"
	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"
)

func Test_BaseURL(t *testing.T) {
	t.Parallel()
	c := NewClient()

	require.NotNil(t, c.BaseURL)
	assert.Equal(t, "api.asset-components.enlight.skf.com", c.BaseURL.Host)

	c = NewClient(WithStage(stages.StageSandbox))

	require.NotNil(t, c.Client.BaseURL)
	assert.Equal(t, "api.sandbox.asset-components.enlight.skf.com", c.BaseURL.Host)
}

func Test_GetComponents(t *testing.T) {
	var (
		ID         = uuid.New()
		Asset      = uuid.New()
		AttachedTo = uuid.New()
		Provider   = uuid.New()
	)
	given := internalmodel.GetAssetComponentsResponse{
		Components: []internalmodel.Component{
			{
				Origin: &internalmodel.Origin{
					Id:       "372812",
					Type:     "@Analyst",
					Provider: string(Provider),
				},
				Manufacturer:        ptr("SKF"),
				AttachedTo:          ptr(string(AttachedTo)),
				Type:                "bearing",
				Asset:               string(Asset),
				PositionDescription: ptr("UNKNOWN"),
				Designation:         ptr("66000"),
				Id:                  ptr(string(ID)),
				SerialNumber:        ptr("123456"),
				RotatingRing:        ptr("inner"),
				ShaftSide:           ptr("de"),
				Position:            1,
				FixedSpeed:          ptr[float32](0),
			},
		},
		Count: ptr[int32](1),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(given)
		require.NoError(t, err)
	}))

	defer server.Close()

	c := NewClient(rest.WithBaseURL(server.URL))

	actual, err := c.GetAssetComponents(context.TODO(), uuid.New(), "test")
	require.NoError(t, err)

	expected := models.Components{
		Components: []models.Component{
			{
				Origin: models.Origin{
					ID:       "372812",
					Type:     "@Analyst",
					Provider: Provider,
				},
				Manufacturer:        "SKF",
				AttachedTo:          AttachedTo,
				Type:                "bearing",
				Asset:               Asset,
				PositionDescription: "UNKNOWN",
				Designation:         "66000",
				ID:                  ID,
				SerialNumber:        "123456",
				RotatingRing:        "inner",
				ShaftSide:           "de",
				Position:            1,
				FixedSpeed:          0,
			},
		},
		Count: 1,
	}
	assert.Equal(t, expected, actual)
}

func ptr[A any](a A) *A {
	return &a
}
