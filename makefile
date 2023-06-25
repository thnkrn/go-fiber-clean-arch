SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage
DOCKERCMD=docker
DOCKER_IMAGE_NAME=go-fiber-clean-arch

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run: ## Start application
	$(GOCMD) run ./cmd/api

test: ## Run tests
	$(GOCMD) test ./... -cover -v

test-coverage: ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out

deps: ## Install dependencies
	$(GOCMD) install github.com/google/wire/cmd/wire@latest
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

deps-cleancache: ## Clear cache in Go module
	$(GOCMD) clean -modcache

wire: ## Generate wire_gen.go
	cd pkg/di && wire

mockery: ## Generate mock package for testing
	cd pkg/usecase && mockery --all --output=../mocks/usecase --case underscore
	cd ../../
	cd pkg/repository && mockery --all --output=../mocks/repository --case underscore

docker-build: ## Build docker image with default setting and platform
	$(DOCKERCMD) build -t $(DOCKER_IMAGE_NAME) .

docker-run: ## Run docker image
	$(DOCKERCMD) run --rm -it -p 8080:8080 $(DOCKER_IMAGE_NAME)

docker-compose-run: ## Run docker image with postgres database in the contianer 
	$(DOCKERCMD) compose up --build       

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'