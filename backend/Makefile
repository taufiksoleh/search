.PHONY: run build clean deps

# Run the server
run:
	go run main.go

# Build the binary
build:
	go build -o bin/server main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run with hot reload (requires air)
dev:
	air
