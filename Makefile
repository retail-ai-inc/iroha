GOPATH=$(shell go env GOPATH)
GOLANGCI_LINT_VERSION=latest

all: lint build

download:
	go mod download
	go mod tidy
	go vet .

build:
	go build -buildvcs=false

build-race: ## build with race detactor
	go build -race -buildvcs=false

build-slim: ## build without symbol and DWARF table, smaller binary but no debugging and profiling ability
	go build -ldflags="-s -w"

lint: ## run all the lint tools, install golangci-lint if not exist
ifeq (,$(wildcard $(GOPATH)/bin/golangci-lint))
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) > /dev/null
	$(GOPATH)/bin/golangci-lint run || exit 0
else
	$(GOPATH)/bin/golangci-lint run || exit 0
endif

clean: ## remove the output binary from go build, as well as go install and build cache
	go clean -i -r -cache

.PHONY: all download build build-race build-slim lint clean
