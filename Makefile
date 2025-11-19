.PHONY: build run stop clean dev prod logs help backend-logs frontend-logs backend frontend install

# Default target
help:
	@echo "AI Search Engine - Available Commands"
	@echo ""
	@echo "Docker Commands:"
	@echo "  make build          - Build all Docker images"
	@echo "  make run            - Run all services with Docker Compose"
	@echo "  make stop           - Stop all running containers"
	@echo "  make prod           - Run with Nginx reverse proxy (production)"
	@echo "  make clean          - Stop and remove all containers, images, volumes"
	@echo "  make logs           - View logs from all containers"
	@echo "  make backend-logs   - View backend container logs"
	@echo "  make frontend-logs  - View frontend container logs"
	@echo ""
	@echo "Development Commands:"
	@echo "  make dev            - Run services in development mode"
	@echo "  make install        - Install dependencies for both projects"
	@echo "  make backend        - Run backend only (local Go)"
	@echo "  make frontend       - Run frontend only (local Node)"
	@echo ""
	@echo "Setup:"
	@echo "  make setup          - Copy .env.example to .env"

# Docker commands
build:
	docker-compose build

run:
	docker-compose up -d

stop:
	docker-compose down

prod:
	docker-compose --profile production up -d --build

clean:
	docker-compose down -v --rmi all --remove-orphans

logs:
	docker-compose logs -f

backend-logs:
	docker-compose logs -f backend

frontend-logs:
	docker-compose logs -f frontend

# Development commands
dev:
	docker-compose up --build

install:
	cd backend && go mod tidy
	cd frontend && npm install

backend:
	cd backend && go run main.go

frontend:
	cd frontend && npm run dev

# Setup
setup:
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "Created .env file. Please add your API keys."; \
	else \
		echo ".env file already exists."; \
	fi

# Rebuild specific service
rebuild-backend:
	docker-compose build backend
	docker-compose up -d backend

rebuild-frontend:
	docker-compose build frontend
	docker-compose up -d frontend

# Health check
health:
	@curl -s http://localhost:8080/api/health | jq . || echo "Backend not running"

# View running containers
ps:
	docker-compose ps
