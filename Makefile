THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build and tag Docker image
	docker build -t gosay:latest .

.PHONY: ci
ci: ## Simulates CI. Runs tests, and builds if they pass
	@$(MAKE) -f $(THIS_FILE) test
	@$(MAKE) -f $(THIS_FILE) build

.PHONY: cleanup
cleanup: ## Stops all containers, and removes temp files created by build/test
	@$(MAKE) -f $(THIS_FILE) stop
	go clean

.PHONY: stop
stop: ## Stop and delete gosay container and images
	docker stop gosay
	docker rm gosay
	docker rmi gosay --force

.PHONY: start
start: ## Stats containers
	docker run -p 8080:8080 --name gosay -d gosay

.PHONY: test
test: ## Runs tests
	go test  -v -p 1 -race -timeout=30s ./...

push: ## push app to docker registry
	docker tag gosay:latest 5587lucas/gosay:latest
	docker push 5587lucas/gosay:latest

.DEFAULT_GOAL := help
