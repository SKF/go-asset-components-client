package main

import (
	"context"
	"fmt"

	components "github.com/SKF/go-asset-components-client/rest"
	"github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"
	dd_http "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

const (
	authToken   = "ACCESS_TOKEN"
	clientID    = "7d77f744-a650-4cc1-a368-2596f7008c68" //random uuid in example
	serviceName = "example-service"
)

func main() {
	ctx := context.Background()

	client := components.NewClient(
		components.WithStage(stages.StageSandbox),
		components.WithClientID(clientID),
		client.WithDatadogTracing(dd_http.RTWithServiceName(serviceName)),
		client.WithTokenProvider(auth.RawToken(authToken)),
	)

	assetID := uuid.New()

	component, err := client.GetComponentsByAsset(ctx, assetID, components.GetComponentsFilter{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", component)
}
