# ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~
#      /\_/\
#     ( o.o )
#      > ^ <
#
# Author: Johan Hanekom
# Date: May 2025
#
# ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~

# =============== // CONSTANTS // ===============

APP_NAME := fortune
SRC := fortune.go
FORTUNES := "fortunes.txt"
BUILD_DIR := build
GO_BUILD := go build

# Targets
.PHONY: all clean build-linux build-windows

# =============== // COMMANDS // ===============

all: build-linux build-windows

build-linux:
	@echo "Building for Linux..."
	cp $(FORTUNES) $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME)-linux $(SRC)

build-windows:
	@echo "Building for Windows..."
	cp $(FORTUNES) $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME)-windows.exe $(SRC)

clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)

