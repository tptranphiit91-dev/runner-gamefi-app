.PHONY: help setup clean get build-runner run test format lint

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

setup: ## Initial setup - install dependencies and generate code
	@echo "📦 Installing dependencies..."
	flutter pub get
	@echo "🔨 Generating code..."
	flutter pub run build_runner build --delete-conflicting-outputs
	@echo "✅ Setup complete!"

clean: ## Clean build artifacts and generated files
	@echo "🧹 Cleaning..."
	flutter clean
	rm -rf .dart_tool
	rm -rf build
	@echo "✅ Clean complete!"

get: ## Get dependencies
	flutter pub get

build-runner: ## Run code generation once
	flutter pub run build_runner build --delete-conflicting-outputs

watch: ## Run code generation in watch mode
	flutter pub run build_runner watch --delete-conflicting-outputs

run: ## Run the app
	flutter run

run-release: ## Run the app in release mode
	flutter run --release

test: ## Run tests
	flutter test

format: ## Format code
	dart format lib/ test/

lint: ## Run linter
	flutter analyze

fix: ## Fix linting issues
	dart fix --apply

build-apk: ## Build Android APK
	flutter build apk --release

build-ios: ## Build iOS app
	flutter build ios --release

rebuild: clean setup ## Clean and rebuild everything
	@echo "✅ Rebuild complete!"

