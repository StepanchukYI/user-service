BIN_NAME := $(or $(PROJECT_NAME),'default-template')
PKG_PATH := $(or $(PKG),'.')
PKG_LIST := $(shell go list ${PKG_PATH}/... | grep -v /vendor/)
GOLINT := golangci-lint

.PHONY: all build clean citest test cover lint check-lint

all: build lint test

dep: # Download required dependencies
	go mod tidy
	go mod download
	go mod vendor

lint: dep check-lint ## Lint the files local env
	$(GOLINT) run --timeout=5m -c .golangci.yml

test: dep ## Run unit tests
	go test -tags=unit -race -count=1 -timeout=5s ./...

cover:
	go test -tags=unit -timeout=5m -cover -v ${PKG_LIST}

msan: ## Run memory sanitizer
	go test -msan -tags=musl,unit ${PKG_LIST}

cilint: dep check-lint
	mkdir reports || true
	$(GOLINT) run --timeout=5m -c .golangci.yml > reports/golangci-lint-report.out
	cat reports/golangci-lint-report.out

citest: dep ## Run tests ci env
	mkdir reports || true
	go test -tags=unit,integration -count=1 -timeout=5m -coverprofile="reports/coverage-report.out" ./... #-race

build: dep ## Build the binary file
	go build -o ./bin/${BIN_NAME} -a .

clean: ## Remove previous build
	rm -f bin/$(BIN_NAME)

fmt: ## format source files
	go fmt ./...

check-lint:
	@which $(GOLINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.25.0

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'