.PHONY: help run build test clean docker-up docker-down migrate

# Variables
APP_NAME=booking-service
MAIN_PATH=cmd/api/main.go
BUILD_DIR=bin

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	@echo "🚀 Starting $(APP_NAME)..."
	go run $(MAIN_PATH)

build: ## Build the application
	@echo "🔨 Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "✅ Build complete: $(BUILD_DIR)/$(APP_NAME)"

test: ## Run tests
	@echo "🧪 Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "🧪 Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report: coverage.html"

clean: ## Clean build artifacts
	@echo "🧹 Cleaning..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	@echo "✅ Clean complete"

deps: ## Download dependencies
	@echo "📦 Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "✅ Dependencies downloaded"

fmt: ## Format code
	@echo "🎨 Formatting code..."
	go fmt ./...
	@echo "✅ Code formatted"

lint: ## Run linter
	@echo "🔍 Running linter..."
	golangci-lint run ./...

docker-up: ## Start PostgreSQL with Docker Compose
	@echo "🐳 Starting Docker containers..."
	docker-compose up -d
	@echo "✅ Docker containers started"

docker-down: ## Stop Docker containers
	@echo "🐳 Stopping Docker containers..."
	docker-compose down
	@echo "✅ Docker containers stopped"

docker-logs: ## Show Docker logs
	docker-compose logs -f

db-create: ## Create database
	@echo "📊 Creating database..."
	createdb booking_db || echo "Database may already exist"
	@echo "✅ Database created"

db-drop: ## Drop database
	@echo "⚠️  Dropping database..."
	dropdb booking_db || echo "Database may not exist"
	@echo "✅ Database dropped"

install-tools: ## Install development tools
	@echo "🔧 Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✅ Tools installed"

.DEFAULT_GOAL := help

