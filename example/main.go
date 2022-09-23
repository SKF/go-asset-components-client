package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"

	components "github.com/SKF/go-asset-component-client/client"
	"github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"
)

const (
	serviceName    = "example-service"
	timeoutSeconds = 10
)

type tokenProvider struct{}

func (t *tokenProvider) GetRawToken(ctx context.Context) (auth.RawToken, error) {
	return auth.RawToken(mustGetEnv("TOKEN")), nil
}

func main() {
	c := components.NewClient(
		components.WithStage(stages.StageSandbox),
		client.WithDatadogTracing(http.RTWithServiceName(serviceName)),
		client.WithTokenProvider(&tokenProvider{}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second) //nolint:
	defer cancel()

	res, err := c.GetAssetComponents(ctx, uuid.New(), "")
	if err != nil {
		panic(err)
	}

	for _, component := range res.Components {
		log.Println(component)
	}
}

func mustGetEnv(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Sprintf("environment variable %q is not set", key))
	}

	return value
}
