# Makefile for macOS Window Manager

# Variables
APP_NAME = window-manager
BINARY_PATH = ./$(APP_NAME)
GO_FILES = $(wildcard *.go)

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build: $(BINARY_PATH)

$(BINARY_PATH): $(GO_FILES)
	@echo "Building $(APP_NAME)..."
	@CGO_ENABLED=1 go build -o $(APP_NAME) .
	@echo "Build complete: $(BINARY_PATH)"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(APP_NAME)
	@echo "Clean complete"

# Run the application in interactive mode
.PHONY: run-interactive
run-interactive: build
	@echo "Starting $(APP_NAME) in interactive mode..."
	@./$(APP_NAME) --interactive

# Run the application
.PHONY: run
run: build
	@echo "Starting $(APP_NAME)..."
	@./$(APP_NAME)

# Test build (compile only, don't run)
.PHONY: test-build
test-build:
	@echo "Testing build..."
	@CGO_ENABLED=1 go build -o /tmp/$(APP_NAME)-test .
	@rm -f /tmp/$(APP_NAME)-test
	@echo "Build test successful"

# Check Go code formatting
.PHONY: fmt
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...

# Vet the code
.PHONY: vet
vet:
	@echo "Vetting Go code..."
	@go vet ./...

# Display help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build      - Build the application"
	@echo "  clean      - Remove build artifacts"
	@echo "  run-interactive - Build and run in interactive mode"
	@echo "  run        - Build and run the application"
	@echo "  test-build - Test that the code compiles"
	@echo "  fmt        - Format Go code"
	@echo "  vet        - Vet Go code"
	@echo "  help       - Show this help message"

# Check if running on macOS
.PHONY: check-platform
check-platform:
	@if [ "$$(uname)" != "Darwin" ]; then \
		echo "Error: This application only works on macOS"; \
		exit 1; \
	fi