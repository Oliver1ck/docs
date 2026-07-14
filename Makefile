BACKEND_DIR := backend

include $(BACKEND_DIR)/.env
export

export ROOT=$(shell pwd)

# --- TG Generator ---
TG_VER ?= 2.3.95
TG_IMAGE := tg:$(TG_VER)
TG_DOCKERFILE := $(BACKEND_DIR)/local/tg/Dockerfile
TG_INPUT ?= $(BACKEND_DIR)/internal/api/interfaces
TG_OUTPUT ?= $(BACKEND_DIR)/internal/api/transport

# Универсальная команда: переменная безвредна на *nix, обязательна на Windows
DOCKER := MSYS_NO_PATHCONV=1 docker

# Создаем алиас для compose с явным указанием путей к конфигам
COMPOSE := $(DOCKER) compose -f $(BACKEND_DIR)/docker-compose.yaml --env-file $(BACKEND_DIR)/.env

.PHONY: tg-build
tg-build:
	@$(DOCKER) build -q -t $(TG_IMAGE) --build-arg TG_VER=$(TG_VER) -f $(TG_DOCKERFILE) .

.PHONY: tg-generate
tg-generate: tg-build
	@mkdir -p $(TG_OUTPUT) $(TG_INPUT)
	@$(DOCKER) run --rm -v "$(ROOT):/app" -w /app $(TG_IMAGE) \
		transport --services $(TG_INPUT) --out $(TG_OUTPUT) --outSwagger $(TG_INPUT)/swagger.yaml
	@echo "Successfully generated. Check $(TG_INPUT)/swagger.yaml"

.PHONY: db-up
db-up:
	@$(COMPOSE) up -d db

.PHONY: db-down
db-down:
	@$(COMPOSE) down db

.PHONY: migrate-up
migrate-up:
	@$(COMPOSE) up --build migrator

.PHONY: migrate-down
migrate-down:
	@$(COMPOSE) run --rm migrator ./migrator-app down

.PHONY: run-backend
run-backend: 
	@cd $(BACKEND_DIR) && go run ./cmd/docs_api

.PHONY: lint
lint: 
	@cd $(BACKEND_DIR) && golangci-lint run

.PHONY: lint-fix
lint-fix: 
	@cd $(BACKEND_DIR) && golangci-lint run --fix