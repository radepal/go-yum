
.DEFAULT_GOAL := help

.PHONY: build
build: ## Make Go build
	go build

.PHONY: help
help: ## List available make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

vet: ## Run vet tool
	go vet .