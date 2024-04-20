VERSION := $(shell git describe --tags --always)

# Targets
.PHONY: all help version
.PHONY: lint

all: help

help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

version: ## show version
	@echo $(VERSION)

lint: ## run golangci-lint
	@golangci-lint run ./...
