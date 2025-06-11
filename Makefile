.DEFAULT_GOAL := help
.PHONY : build

MAKEFILE := $(abspath $(firstword $(MAKEFILE_LIST)))
MAKEFILE_DIR := $(abspath $(dir $(MAKEFILE)))
BIN_DIR := $(MAKEFILE_DIR)/bin

GOPATH := $(shell go env GOPATH)
DESTDIR?=
PREFIX?=/usr/local
ENV?=development

VERSION := $(shell git describe --always)

RFC_3339 := "+%Y-%m-%dT%H:%M:%SZ"
DATE := $(shell date -u $(RFC_3339))
COMMIT := $(shell git rev-list -1 HEAD)

OPTS?=GO111MODULE=on

SWAGGER_PORT?=8001
SWAGGER_MODE?=editor

# Database connection string
APP :=finance_tracker
DB_CS="host=localhost port=5432 user=tejashree password=password dbname=$(APP) sslmode=disable"
ifeq ($(ENV),review)
	DB_CS="host=nonprod-readonly-db.cafu.app port=5432 user=postgres password= dbname=$(APP)_review sslmode=disable"
endif

OPENAPI_DOC := "./docs/api-merge.yml"

run: ## Run code
	@go run ./cmd/server/

build: ## Build binary
	@mkdir -p bin
	@go build  -o bin/$(APP) ./cmd/server/

test: ## Run tests
	@go test ./... -race -coverprofile=./.coverage/coverage.out -covermode=atomic -coverpkg=./...

push-checks: build test lint client speccy

coverage: ## show coverage
	@go tool cover -html=./.coverage/coverage.out

cover: ## total coverage in cmd
	@make test
	@go tool cover -func=./.coverage/coverage.out

lint: ## Run linters
	@golangci-lint run --timeout=3m --issues-exit-code=0

lint-html: ## Run linters and output html format
	@golangci-lint run --issues-exit-code 0 --out-format html > gl-code-quality-report.html

clean: ## Cleaning binary
	-rm -f bin/$(APP)
	-rm ./.coverage/coverage.out

docker: ## Start all dependent containers and migrate DB schema
	cd ./build && docker compose up -d
	sleep 8 ## wait till Postgres is ready :)
	make migration-up

down: ## Stops all dependent containers
	-cd ./build && docker compose down

reset: clean down ## Remove all dependent containers with data and start again
	-rm -rf ./build/postgres/*
	make swagger-reset
	make docker

migration-status: ## Migration status
	@goose -dir=./database/migrations postgres $(DB_CS) status

migration-up: ## Migration up
	@goose -dir=./database/migrations postgres $(DB_CS) up

migration-down: ## Migration down
	@goose -dir=./database/migrations postgres $(DB_CS) down

migration-create: ## Migration create (migration-create init)
	goose -dir=./database/migrations postgres $(DB_CS) create $(*) sql

swagger-up: build-api ## Swagger editor up
	@echo "dir:" $(PWD)
	@docker run -d -p $(SWAGGER_PORT):8080 --platform linux/amd64 --name swagger_editor -v $(PWD)/./docs/:/tmp/docs -e SWAGGER_FILE=/tmp/docs/api-merge.yml swaggerapi/swagger-editor
	@echo "Running swagger at http://localhost:$(SWAGGER_PORT)"

swagger-down: ## Swagger editor down
	-docker kill swagger_editor
	-docker rm swagger_editor

swagger-reset: swagger-down swagger-up

swagger-watcher:
	swagger-ui-watcher ./docs/api-merge.yml -p 8080

speccy:
	speccy lint ./docs/api.yml

build-api:
	@redocly bundle ./docs/api.yml --ext=yml --output=./docs/api-merge.yml

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

## disclaimer: this is not a full automated feature and should be used carefully 
## https://github.com/OpenAPITools/openapi-generator
generate: build-api ## Generate models from OpenAPI documentation and move it to controllers
	@openapi-generator generate --global-property models -i $(OPENAPI_DOC) -g go-gin-server -o ./generated/openapi/
	@cd ./generated/openapi/go/ && sed -i '.bak' 's/package openapi/package model/g' * && rm *.bak
	@mv ./generated/openapi/go/model*.go ./internal/app/controller/model/
	@cd ./internal/app/controller/model/ && rename -f "s/model_*//" *
	@rm -rf ./generated/openapi/go
	@openapi-generator generate -i $(OPENAPI_DOC) -g go -o ../generated/client/openapi

spell: ## spell checking
	misspell -i importas .
	identypo ./...

install-tools: ## Install / Update tools which are required to run Makefile commands
	npm install speccy -g
	brew install openapi-generator
	brew install goose
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/alexkohler/identypo/cmd/identypo@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.0
	# Install redocly to merge API doc
	@npm i -g @redocly/cli@latest

install-swagger-tools: ## Install dev tools
	# Install swaggo for API doc generation
	@GOBIN="$(BIN_DIR)" go install github.com/swaggo/swag/cmd/swag@latest
	# Install redocly CLI for API doc manipulation
	@rm -rf /opt/homebrew/lib/node_modules/@redocly
	@npm install -g @redocly/cli || sudo npm install -g @redocly/cli
	# Install speccy for linting API doc
	@rm -rf /opt/homebrew/lib/node_modules/speccy
	@npm install -g speccy || sudo npm install -g speccy
	# Install swagger-ui-watcher to reload Swagger UI in your browser

swagger-gen: ## Generate API doc
	# Remove all temp files that might be there because of a previously failed doc-gen.
	@rm -rf ./docs/tmp
	# Generate OpenAPI v2 doc from swaggo/swag annotations.
	@$(BIN_DIR)/swag init -g ./cmd/server/main.go -o ./docs/tmp --parseDependency --parseInternal --quiet --collectionFormat multi
	# Convert the generated OpenAPI v2 yaml file to OpenAPI v3 yaml file.
	@docker run --rm -u $(shell id -u):$(shell id -g) -v $(PWD)/docs/tmp:/work openapitools/openapi-generator-cli:latest-release \
        generate -i /work/swagger.yaml -o /work/v3 -g openapi-yaml --minimal-update 1> /dev/null
	# Remove the path prefix from the generated schema names.
	@sed -i -e "s/gitlab_intelligentb_com_.*\.//g" ./docs/tmp/v3/openapi/openapi.yaml
	@sed -i -e "s/gitlab_intelligentb_com_.*_models_//g" ./docs/tmp/v3/openapi/openapi.yaml
	@for prefix in types apierror model models pagination time; do \
		sed -i -e "s/$$prefix\.//g" ./docs/tmp/v3/openapi/openapi.yaml ; \
    done
	@sed -i -e "s/type\: object//g" ./docs/tmp/v3/openapi/openapi.yaml
	@sleep 1
	# Replace the servers section of the above temp file. This is because swaggo/swag only supports OpenAPI v2 and v2 doesn't support multiple servers.
	@docker run --security-opt=no-new-privileges --cap-drop all --network none --rm -v $(PWD)/docs:/work mikefarah/yq '. *n load("/work/tmp/v3/openapi/openapi.yaml")' /work/overriding-template.yml > ./docs/api.yml
	@sleep 1
	# Remove all temp files.
	@rm -rf ./docs/tmp
	# Check the final API doc.
	@speccy lint -v ./docs/api.yml
