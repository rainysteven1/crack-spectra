# Global Variables
BACKEND_DIR = backend
NS          = crack-spectra

.PHONY: help build run debug swagger import clean-finalizer delete-chaos k8s-resources ports \
        install-hooks git-sync upgrade-dep deploy-ts

.DEFAULT_GOAL := help

help:  ## Display targets with category headers
	@awk 'BEGIN { \
		FS = ":.*##"; \
		printf "\n\033[1;34mUsage:\033[0m\n  make \033[36m<target>\033[0m\n\n\033[1;34mTargets:\033[0m\n"; \
	} \
	/^##@/ { \
		header = substr($$0, 5); \
		printf "\n\033[1;33m%s\033[0m\n", header; \
	} \
	/^[a-zA-Z_-]+:.*?##/ { \
		printf "  \033[36m%-20s\033[0m \033[90m%s\033[0m\n", $$1, $$2; \
	}' $(MAKEFILE_LIST)

##@ Building

run: ## Build and deploy using skaffold
	skaffold run --default-repo=$(DEFAULT_REPO)

##@ Development

local-debug: ## Start local debug environment (databases + controller)
	docker compose down && \
	docker compose up redis mariadb -d && \
	cd $(BACKEND_DIR) && go run main.go both --conf config.dev.toml --port 8082

swagger: ## Generate Swagger API documentation
	swag init -d ./$(BACKEND_DIR) --parseDependency --parseDepth 1 --output ./$(BACKEND_DIR)/docs

##@ Git Management

install-hooks: ## Install pre-commit hooks
	chmod +x scripts/hooks/pre-commit
	cp scripts/hooks/pre-commit .git/hooks/pre-commit

git-sync: ## Synchronize Git submodules
	git submodule update --init --recursive --remote

upgrade-dep: git-sync ## Upgrade Git submodules to latest main branch
	@git submodule foreach 'branch=$$(git config -f $$toplevel/.gitmodules submodule.$$name.branch || echo main); \
		echo "Updating $$name to branch: $$branch"; \
		git checkout $$branch && git pull origin $$branch'
