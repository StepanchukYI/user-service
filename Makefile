BIN_NAME := $(or $(PROJECT_NAME),main)
.PHONY: all build clean test cover

all: build test

dep: # Download required dependencies
	go mod tidy
	go mod download
	go mod vendor

test: dep ## Run unit tests
	go test -race -count=1 -timeout=5s -v ./...

cover:
	go test -timeout=5m -cover -v ./...

build: dep ## Build the binary file
	CGO_ENABLED=0 xGOARCH=amd64 GOHOSTARCH=amd64 GOHOSTOS=linux GOOS=linux go build -o ./bin/${BIN_NAME} -a .

clean: ## Remove previous build
	rm -f bin/$(BIN_NAME)

dcb:
	docker-compose -f ./deployments/docker-compose.yaml build --no-cache

dcu:
	docker-compose -f ./deployments/docker-compose.yaml up
