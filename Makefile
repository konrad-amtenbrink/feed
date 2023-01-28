build: ## Build the application
	@go build -o app main.go 

lint: ## Lint the application
	@golangci-lint run

setup: ## Setup the local dev environment
	@docker compose -f .development/docker-compose.yml up -d

run: ## Run the application
	@go run main.go server --port 4201

mock: ## Create the database and storage mocks
	mockgen -source=db/db.go -destination=db/db_mock.go && mockgen -source=storage/storage.go -destination=storage/storage_mock.go

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
