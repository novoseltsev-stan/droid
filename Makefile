# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


lint: ## Run linters
	@golangci-lint run

fix: ## Autofix linter issues
	@golangci-lint run --fix

fmt: ## Format code
	@golangci-lint fmt



test: ## Run tests
	@go test ./... -tags=integration -v -json -race  -coverprofile=.reports/profile.out -covermode=atomic | go tool tparse -all

test-fast: ## Run fast unit tests
	@go test  ./...  -v -json | go tool tparse -all


update: ## Update dependencies
	@go get -u -t ./... && go get tool && go mod tidy

gen: ## Generate code
	@go generate ./...

build: ## Build cli
	@go build -a -ldflags="-w -s" -o ./droid ./cmd/droid

