WGET   ?= wget
DOCKER ?= docker

API_URL := "https://api.sandbox.asset-components.enlight.skf.com"
OPENAPITOOLS_VERSION := v6.1.0

.PHONY: all
all: rest/models/ rest/openapi.yaml

rest/openapi.yaml:
	$(WGET) "$(API_URL)/docs/swagger/openapi.yaml" -O "$@"

rest/models/: rest/openapi.yaml
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
			--output $@

.PHONY: clean
clean:
	$(RM) rest/models/model_* rest/openapi.yaml