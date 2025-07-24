.PHONY: build run test clean docker-up docker-down migrate

# Build the application
build:
	go build -o bin/main cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Run tests
test:
	go test -v ./tests/...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./tests/...
	go tool cover -html=coverage.out

# Clean build files
clean:
	rm -rf bin/

# Docker commands
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-build:
	docker build -t github.com/faizalnurrozi/go-starter-kit .

# Database migration
migrate-up:
	migrate -path migrations -database "mysql://root:@localhost:3306/go_base_project_db?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "mysql://root:@localhost:3306/go_base_project_db?sslmode=disable" down

# Generate gRPC code
gen-proto:
	protoc --go_out=. --go-grpc_out=. proto/user/user.proto

# Install dependencies
deps:
	go mod tidy
	go mod download

# Lint code
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Development setup
dev-setup: deps docker-up
	@echo "Development environment is ready!"

# Production build
prod-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/main cmd/server/main.go
