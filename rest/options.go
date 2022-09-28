package rest

import (
	"fmt"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
)

func WithStage(stage string) Option {
	if stage == stages.StageProd {
		return rest.WithBaseURL("https://api.asset-components.enlight.skf.com/")
	}

	return rest.WithBaseURL(fmt.Sprintf("https://api.%s.asset-components.enlight.skf.com/", stage))
}

func WithClientID(clientID string) Option {
	return rest.WithDefaultHeader("X-Client-Id", clientID)
}
