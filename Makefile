# Makefile for encrypt-conn-tools
# Requires garble: go install mvdan.cc/garble@latest
# Requires cross-compilers (CC) for CGO builds if building on a different host OS.

APP_NAME := libencrypt
OUTPUT_DIR := exe
SRC := main.go
GARBLE_FLAGS := -tiny -literals -seed=random

# Default target
.PHONY: all
all: linux-amd64 linux-arm64 windows-amd64 windows-arm64 darwin-amd64 darwin-arm64

# Create output directory
$(OUTPUT_DIR):
	mkdir -p $(OUTPUT_DIR)

# Linux AMD64
.PHONY: linux-amd64
linux-amd64: $(OUTPUT_DIR)
	@echo "Building for Linux AMD64..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 \
	garble $(GARBLE_FLAGS) build -buildmode=c-shared -o $(OUTPUT_DIR)/$(APP_NAME)-linux-amd64.so $(SRC)

# Linux ARM64
.PHONY: linux-arm64
linux-arm64: $(OUTPUT_DIR)
	@echo "Building for Linux ARM64..."
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 \
	garble $(GARBLE_FLAGS) build -buildmode=c-shared -o $(OUTPUT_DIR)/$(APP_NAME)-linux-arm64.so $(SRC)

# Windows AMD64
.PHONY: windows-amd64
windows-amd64: $(OUTPUT_DIR)
	@echo "Building for Windows AMD64..."
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
	garble $(GARBLE_FLAGS) build -buildmode=c-shared -o $(OUTPUT_DIR)/$(APP_NAME)-windows-amd64.dll $(SRC)

# Windows ARM64
.PHONY: windows-arm64
windows-arm64: $(OUTPUT_DIR)
	@echo "Building for Windows ARM64..."
	GOOS=windows GOARCH=arm64 CGO_ENABLED=1 \
	garble $(GARBLE_FLAGS) build -buildmode=c-shared -o $(OUTPUT_DIR)/$(APP_NAME)-windows-arm64.dll $(SRC)

# macOS AMD64 (Intel)
.PHONY: darwin-amd64
darwin-amd64: $(OUTPUT_DIR)
	@echo "Building for macOS AMD64..."
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 \
	garble $(GARBLE_FLAGS) build -buildmode=c-shared -o $(OUTPUT_DIR)/$(APP_NAME)-darwin-amd64.dylib $(SRC)

# macOS ARM64 (Apple Silicon)
.PHONY: darwin-arm64
darwin-arm64: $(OUTPUT_DIR)
	@echo "Building for macOS ARM64..."
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 \
	garble $(GARBLE_FLAGS) build -buildmode=c-shared -o $(OUTPUT_DIR)/$(APP_NAME)-darwin-arm64.dylib $(SRC)

# Clean
.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR)

