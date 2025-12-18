.PHONY: help wire build run test clean
.PHONY: loc-infra loc-backend loc-all
.PHONY: dev-infra dev-backend dev-all
.PHONY: prod-infra prod-backend prod-all
.PHONY: down logs

# ==================== Help ====================

help:
	@echo "SyncDrive Backend - Available Commands"
	@echo ""
	@echo "Development:"
	@echo "  make wire           - Generate Wire dependency injection code"
	@echo "  make build          - Build application binary"
	@echo "  make run            - Run API server directly (ENV=loc)"
	@echo "  make test           - Run tests"
	@echo "  make clean          - Clean build artifacts and logs"
	@echo ""
	@echo "Local Development (loc):"
	@echo "  make loc-infra      - Start infrastructure only"
	@echo "  make loc-backend    - Run API server with go run"
	@echo "  make loc-all        - Start all services with docker-compose"
	@echo ""
	@echo "Development Environment (dev):"
	@echo "  make dev-infra      - Start infrastructure"
	@echo "  make dev-backend    - Start API with docker-compose"
	@echo "  make dev-all        - Start all services"
	@echo ""
	@echo "Production Environment (prod):"
	@echo "  make prod-infra     - Start infrastructure"
	@echo "  make prod-backend   - Start API with docker-compose"
	@echo "  make prod-all       - Start all services"
	@echo ""
	@echo "Utilities:"
	@echo "  make down           - Stop all docker services"
	@echo "  make logs           - View docker logs"

# ==================== Development ====================

wire:
	@if [ ! -f $(HOME)/go/bin/wire ]; then \
		echo "Wire not found, installing..."; \
		go install github.com/google/wire/cmd/wire@latest; \
	fi
	cd cmd/api && $(HOME)/go/bin/wire

build:
	mkdir -p bin
	go build -ldflags="-w -s" -o bin/api ./cmd/api/

run:
	ENV=loc go run ./cmd/api/

test:
	go test -v -race ./...

clean:
	rm -rf bin/
	rm -rf logs/*

# ==================== Local Development (loc) ====================

loc-infra:
	cd deployments && docker-compose --profile infra up -d

loc-backend:
	ENV=loc go run ./cmd/api/

loc-all:
	cd deployments && docker-compose --profile infra --profile backend up -d

# ==================== Development Environment (dev) ====================

dev-infra:
	cd deployments && ENV=dev docker-compose --profile infra up -d

dev-backend:
	cd deployments && ENV=dev GIN_MODE=debug docker-compose --profile backend up -d

dev-all:
	cd deployments && ENV=dev GIN_MODE=debug docker-compose --profile infra --profile backend up -d

# ==================== Production Environment (prod) ====================

prod-infra:
	cd deployments && ENV=prod docker-compose --profile infra up -d

prod-backend:
	cd deployments && ENV=prod GIN_MODE=release docker-compose --profile backend up -d

prod-all:
	cd deployments && ENV=prod GIN_MODE=release docker-compose --profile infra --profile backend up -d

# ==================== Utilities ====================

down:
	cd deployments && docker-compose --profile infra --profile backend down

logs:
	cd deployments && docker-compose logs -f
