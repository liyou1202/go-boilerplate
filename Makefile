.PHONY: help wire test clean down logs
.PHONY: loc-backend loc-infra loc-all
.PHONY: dev-backend dev-infra dev-all
.PHONY: prod-backend prod-infra prod-all

# ==================== Help ====================

help:
	@echo "SyncDrive Backend - Available Commands"
	@echo ""
	@echo "Local Development (loc):"
	@echo "  make loc-backend    - Run backend API with go run (ENV=loc)"
	@echo "  make loc-infra      - Start infrastructure services only"
	@echo "  make loc-all        - Start infrastructure + backend"
	@echo ""
	@echo "Development Environment (dev):"
	@echo "  make dev-backend    - Start backend API with docker-compose (ENV=dev)"
	@echo "  make dev-infra      - Start infrastructure services only"
	@echo "  make dev-all        - Start infrastructure + backend"
	@echo ""
	@echo "Production Environment (prod):"
	@echo "  make prod-backend   - Start backend API with docker-compose (ENV=prod)"
	@echo "  make prod-infra     - Start infrastructure services only"
	@echo "  make prod-all       - Start infrastructure + backend"
	@echo ""
	@echo "Development Tools:"
	@echo "  make wire           - Generate Wire dependency injection code"
	@echo "  make test           - Run tests"
	@echo "  make clean          - Clean build artifacts and logs"
	@echo ""
	@echo "Docker Utilities:"
	@echo "  make down           - Stop all docker services"
	@echo "  make logs           - View docker logs"

# ==================== Development Tools ====================

wire:
	@if [ ! -f $(HOME)/go/bin/wire ]; then \
		echo "Wire not found, installing..."; \
		go install github.com/google/wire/cmd/wire@latest; \
	fi
	cd cmd/api && $(HOME)/go/bin/wire

test:
	go test -v -race ./...

clean:
	rm -rf bin/
	rm -rf logs/*

# ==================== Local Development (loc) ====================

loc-backend:
	make wire
	ENV=loc GIN_MODE=debug go run ./cmd/api/

loc-infra:
	cd deployments && docker-compose --profile infra up -d

loc-all:
	make wire
	cd deployments && docker-compose --profile infra --profile backend up -d

# ==================== Development Environment (dev) ====================

dev-backend:
	make wire
	cd deployments && ENV=dev GIN_MODE=debug docker-compose --profile backend up -d

dev-infra:
	cd deployments && ENV=dev docker-compose --profile infra up -d

dev-all:
	make wire
	cd deployments && ENV=dev GIN_MODE=debug docker-compose --profile infra --profile backend up -d

# ==================== Production Environment (prod) ====================

prod-backend:
	make wire
	cd deployments && ENV=prod GIN_MODE=release docker-compose --profile backend up -d

prod-infra:
	cd deployments && ENV=prod docker-compose --profile infra up -d

prod-all:
	make wire
	cd deployments && ENV=prod GIN_MODE=release docker-compose --profile infra --profile backend up -d

# ==================== Docker Utilities ====================

down:
	cd deployments && docker-compose --profile infra --profile backend down

logs:
	cd deployments && docker-compose logs -f
