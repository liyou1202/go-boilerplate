.PHONY: help build test clean
.PHONY: loc-infra loc-backend loc-all
.PHONY: dev-infra dev-backend dev-all
.PHONY: prod-infra prod-backend prod-all
.PHONY: down logs

COMPOSE_FILE := deployments/docker-compose.yml

help:
	@echo "SyncDrive Backend - Available Commands"
	@echo ""
	@echo "Local Development (loc):"
	@echo "  make loc-infra      - Start infrastructure only (MySQL, MongoDB, Redis, MQTT)"
	@echo "  make loc-backend    - Run API server locally (go run)"
	@echo "  make loc-all        - Start all services with docker-compose"
	@echo ""
	@echo "Development Environment (dev):"
	@echo "  make dev-infra      - Start infrastructure only"
	@echo "  make dev-backend    - Start API with docker-compose"
	@echo "  make dev-all        - Start all services"
	@echo ""
	@echo "Production Environment (prod):"
	@echo "  make prod-infra     - Start infrastructure only"
	@echo "  make prod-backend   - Start API with docker-compose"
	@echo "  make prod-all       - Start all services"
	@echo ""
	@echo "Utilities:"
	@echo "  make build          - Build application binary"
	@echo "  make test           - Run tests"
	@echo "  make down           - Stop all docker services"
	@echo "  make logs           - View docker logs"
	@echo "  make clean          - Clean build artifacts"

# ==================== Local Development (loc) ====================

loc-infra:
	@echo "Starting infrastructure (loc)..."
	@cd deployments && docker-compose --profile infra up -d
	@echo "Infrastructure started. Services available at localhost"

loc-backend:
	@echo "Running API server locally (ENV=loc)..."
	@ENV=loc go run cmd/api/main.go

loc-all:
	@echo "Starting all services (loc)..."
	@cd deployments && docker-compose --profile infra --profile backend up -d
	@echo "All services started"

# ==================== Development Environment (dev) ====================

dev-infra:
	@echo "Starting infrastructure (dev)..."
	@cd deployments && ENV=dev docker-compose --profile infra up -d
	@echo "Infrastructure started"

dev-backend:
	@echo "Starting API server (dev)..."
	@cd deployments && ENV=dev GIN_MODE=debug docker-compose --profile backend up -d
	@echo "API server started"

dev-all:
	@echo "Starting all services (dev)..."
	@cd deployments && ENV=dev GIN_MODE=debug docker-compose --profile infra --profile backend up -d
	@echo "All services started"

# ==================== Production Environment (prod) ====================

prod-infra:
	@echo "Starting infrastructure (prod)..."
	@cd deployments && ENV=prod docker-compose --profile infra up -d
	@echo "Infrastructure started"

prod-backend:
	@echo "Starting API server (prod)..."
	@cd deployments && ENV=prod GIN_MODE=release docker-compose --profile backend up -d
	@echo "API server started"

prod-all:
	@echo "Starting all services (prod)..."
	@cd deployments && ENV=prod GIN_MODE=release docker-compose --profile infra --profile backend up -d
	@echo "All services started"

# ==================== Utilities ====================

build:
	@mkdir -p bin
	@go build -ldflags="-w -s" -o bin/api cmd/api/main.go
	@echo "Build complete: bin/api"

test:
	@go test -v -race ./...

down:
	@echo "Stopping all services..."
	@cd deployments && docker-compose --profile infra --profile backend down
	@echo "Services stopped"

logs:
	@cd deployments && docker-compose logs -f

clean:
	@rm -rf bin/ logs/*
	@echo "Clean complete"
