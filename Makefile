# Makefile for bunchotext

BINARY := bunchotext
CMD_DIR := ./cmd/bunchotext
VERSION := 1.0.0

# Platform-specific settings
ifeq ($(OS),Windows_NT)
    INSTALL_DIR := $(HOME)/bin/$(BINARY).exe
    BINARY := $(BINARY).exe
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Darwin)
        INSTALL_DIR := $(HOME)/.local/bin/$(BINARY)
    else
        INSTALL_DIR := $(HOME)/.local/bin/$(BINARY)
    endif
endif

.PHONY: all build clean install release help

all: build

## build: Build the binary for current platform
build:
	@echo "Building bunchotext..."
	go build -o $(BINARY) $(CMD_DIR)
	@echo "Build complete: $(BINARY)"

## build-linux: Build for Linux (amd64)
build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY)-linux-amd64 $(CMD_DIR)

## build-mac: Build for macOS (amd64 & arm64)
build-mac:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY)-darwin-amd64 $(CMD_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY)-darwin-arm64 $(CMD_DIR)

## build-windows: Build for Windows (amd64)
build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(BINARY)-windows-amd64.exe $(CMD_DIR)

## release: Build all platforms for release
release: clean build-linux build-mac build-windows
	@echo "All release builds complete!"
	@ls -lh $(BINARY)-*

## install: Build and install to local bin directory
install: build
	@echo "Installing to $(INSTALL_DIR)..."
	@mkdir -p $(HOME)/.local/bin
	@cp $(BINARY) $(INSTALL_DIR)
	@chmod +x $(INSTALL_DIR)
	@echo "Installation complete! Run '$(BINARY)' from anywhere."

## uninstall: Remove installed binary
uninstall:
	@echo "Removing $(INSTALL_DIR)..."
	@rm -f $(INSTALL_DIR)
	@echo "Uninstallation complete."

## clean: Remove build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(BINARY) $(BINARY)-* *.exe
	@echo "Clean complete."

## help: Show this help message
help:
	@echo "bunchotext Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build          - Build for current platform"
	@echo "  make build-linux    - Build for Linux"
	@echo "  make build-mac      - Build for macOS"
	@echo "  make build-windows  - Build for Windows"
	@echo "  make release        - Build all platforms"
	@echo "  make install        - Build and install locally"
	@echo "  make uninstall      - Remove installed binary"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make help           - Show this help"