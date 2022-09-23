.PHONY: clean internal/models
.INTERMEDIATE: swagger.yaml

RM     ?= rm
WGET   ?= wget
MKDIR  ?= mkdir
DOCKER ?= docker

API_URL := "https://api.sandbox.asset-components.enlight.skf.com"
OPENAPITOOLS_VERSION := v6.1.0

swagger.yaml:
	$(WGET) "$(API_URL)/docs/swagger/openapi.yaml" -O "$@"

internal/models: swagger.yaml
	$(RM) -rf "$@" && $(MKDIR) -p "$@"
	$(DOCKER) run --rm \
		--volume $(shell pwd):/src \
		--workdir /src \
		--user "$(shell id -u):$(shell id -g)" \
		openapitools/openapi-generator-cli:${OPENAPITOOLS_VERSION} \
			generate \
			--input-spec $< \
			--global-property models,modelDocs=false \
			--generator-name go \
			--additional-properties packageName=models \
			--output internal/models