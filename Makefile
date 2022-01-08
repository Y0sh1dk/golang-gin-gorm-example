# Makefile

.PHONY: help
help: ## Show help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Starts the stack
	docker-compose --env-file ./config/.env up

.PHONY: run-background
run-background: ## Starts the stack in the background
	docker-compose --env-file ./config/.env up -d

.PHONY: stop
stop: ## Stops the stack
	docker-compose stop

.PHONY: clean
clean: ## Cleans up
	rm -rf ./tmp
