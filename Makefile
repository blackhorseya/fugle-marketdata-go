# Variables
GO := go
BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin
LDFLAGS := -w -s

VERSION := $(shell git describe --tags --always)

# Targets
.PHONY: all help version
.PHONY: lint clean

all: help

help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

version: ## show version
	@echo $(VERSION)

clean: ## clean build directory
	@rm -rf $(BUILD_DIR)

## go
lint: ## run golangci-lint
	@golangci-lint run ./...

build: ## build binary
	$(GO) build -ldflags "$(LDFLAGS)" ./...

test: ## run tests
	$(GO) test -v ./...
