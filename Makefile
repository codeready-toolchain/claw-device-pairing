.PHONY: build test run lint

.DEFAULT_GOAL := help

# Image tag (can be overridden with IMG variable)
IMG ?= quay.io/xcoulon/claw-device-pairing:latest

PLATFORM ?= linux/amd64

# Version information
COMMIT_HASH := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS := -X 'github.com/xcoulon/claw-device-pairing/internal/version.CommitHash=$(COMMIT_HASH)' \
           -X 'github.com/xcoulon/claw-device-pairing/internal/version.BuildTime=$(BUILD_TIME)'

help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build-server: ## Build the server binary to bin/claw-device-pairing
	@echo "Building claw-device-pairing (commit: $(COMMIT_HASH), build time: $(BUILD_TIME))..."
	@mkdir -p bin
	go build -ldflags "$(LDFLAGS)" -o bin/claw-device-pairing cmd/main.go
	@echo "Build complete: bin/claw-device-pairing"

test: ## Run Go tests with coverage
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...
	@echo "Test complete. Coverage report: coverage.out"

run-server: build-server ## Build and run the server on port 8080
	@echo "Starting server..."
	./bin/claw-device-pairing serve

run-ui: ## Run the frontend development server
	@echo "Starting frontend dev server..."
	cd ui && npm run dev

clean: ## Remove build artifacts (bin/ directory)
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	@echo "Clean complete"

build-image: ## Build container image using Podman (use IMG variable to customize tag)
	@echo "Building container image: $(IMG) (commit: $(COMMIT_HASH), build time: $(BUILD_TIME))"
	podman build --platform=$(PLATFORM) \
		--build-arg COMMIT_HASH=$(COMMIT_HASH) \
		--build-arg BUILD_TIME=$(BUILD_TIME) \
		-t $(IMG) -f Containerfile .
	@echo "Image built: $(IMG)"

push-image: ## Push container image to registry (use IMG variable to customize tag)
	@echo "Pushing image: $(IMG)"
	podman push $(IMG)
	@echo "Image pushed: $(IMG)"