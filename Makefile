.PHONY: test coverage lint fmt vet build clean examples

# Run tests
test:
	go test -v ./...

# Run tests with coverage
coverage:
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Build examples
build:
	go build -o bin/basic examples/basic/main.go
	go build -o bin/manual examples/manual/main.go

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.txt coverage.html

# Run examples
examples: build
	@echo "Running basic example..."
	./bin/basic
	@echo "\nRunning manual example..."
	./bin/manual

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run all checks
check: fmt vet test

.DEFAULT_GOAL := test
