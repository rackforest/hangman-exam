# Variables
BINARY_NAME=hangman
SRC=$(wildcard *.go)

# Build the application
build:
	go build -o $(BINARY_NAME) $(SRC)

# Run the application
run: build
	./$(BINARY_NAME)

# Test the application
test:
	go test ./...

# Clean up build artifacts
clean:
	rm -f $(BINARY_NAME)

# Phony targets
.PHONY: build run test clean
