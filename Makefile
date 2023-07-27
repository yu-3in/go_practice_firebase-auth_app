.PHONY: help build build-local up down logs ps test \
dry-migrate migrate generate
.DEFAULT_GOAL := help

DOCKER_TAG := latest

up: ## Start the project
	docker-compose up -d

down: ## Stop the project
	docker-compose down

build: ## Build docker image to production
	docker build --target production ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	go test -race -shuffle=on ./...

generate: ## Generate codes
	go generate ./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
