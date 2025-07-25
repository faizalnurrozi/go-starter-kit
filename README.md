## README.md

# Go Base Project

Robust Go application with Fiber framework implementing clean architecture and SOLID principles.

## Features

- **Clean Architecture**: Separation of concerns with layers (Handler, Service, Repository)
- **SOLID Principles**: Dependency injection, interface segregation
- **Database Support**: MySQL, PostgreSQL with GORM
- **Caching**: Redis integration
- **gRPC Support**: Protocol buffer definitions and handlers
- **API Versioning**: v1, v2, etc. with proper routing
- **Middleware**: Authentication, logging, CORS, validation
- **Centralized Logging**: Structured logging with Logrus
- **Error Handling**: Standardized error responses
- **Validation**: Request/parameter validation
- **Testing**: Unit and integration tests
- **Docker**: Ready-to-use Docker setup
- **Configuration**: Environment-based config with Viper

## Project Structure

```
project/
├── cmd/server/          # Application entry point
├── internal/            # Private application code
│   ├── config/          # Configuration management
│   ├── database/        # Database connection and migration
│   ├── cache/           # Redis cache implementation
│   ├── grpc/            # gRPC server and handlers
│   ├── middleware/      # HTTP middleware
│   ├── handler/         # HTTP handlers (controllers)
│   ├── service/         # Business logic layer
│   ├── repository/      # Data access layer
│   ├── dto/             # Data transfer objects
│   ├── entity/          # Database entities
│   ├── logger/          # Centralized logging
│   ├── errors/          # Custom error definitions
│   └── utils/           # Utility functions
├── proto/               # gRPC protocol definitions
├── tests/               # Test files
│   ├── unit/            # Unit tests
│   └── integration/     # Integration tests
├── migrations/          # Database migrations
├── pkg/                 # Public packages
└── configs/             # Configuration files
```

## Quick Start

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- PostgreSQL/MySQL
- Redis

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd github.com/faizalnurrozi/go-starter-kit
```

2. Install dependencies:
```bash
make deps
```

3. Copy environment file:
```bash
cp .env.example .env
```

4. Start dependencies with Docker:
```bash
make docker-up
```

5. Run migrations:
```bash
make migrate-up
```

6. Start the application:
```bash
make run
```

The application will be available at:
- HTTP API: `http://localhost:8080`
- gRPC: `localhost:9090`
- Health Check: `http://localhost:8080/health`

## API Documentation

### Authentication

All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

### User Endpoints

#### Create User
```http
POST /api/v1/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Get User by ID
```http
GET /api/v1/users/{id}
Authorization: Bearer <token>
```

#### Get All Users
```http
GET /api/v1/users?limit=10&offset=0
Authorization: Bearer <token>
```

#### Update User
```http
PUT /api/v1/users/{id}
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "is_active": true
}
```

#### Delete User
```http
DELETE /api/v1/users/{id}
Authorization: Bearer <token>
```

### Response Format

All API responses follow this standard format:

#### Success Response (200)
```json
{
  "status": "success",
  "code": 200,
  "message": "Request successful",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Business Error (404)
```json
{
  "status": "error",
  "code": 404,
  "message": "User not found",
  "error": "Additional error details"
}
```

#### System Error (500)
```json
{
  "status": "error",
  "code": 500,
  "message": "Internal server error"
}
```

#### Validation Error (400)
```json
{
  "status": "error",
  "code": 400,
  "message": "Validation failed: name required, email invalid"
}
```

#### Unauthorized (401)
```json
{
  "status": "error",
  "code": 401,
  "message": "Unauthorized"
}
```

## Testing

### Run All Tests
```bash
make test
```

### Run Tests with Coverage
```bash
make test-coverage
```

### Unit Tests
```bash
go test -v ./tests/unit/...
```

### Integration Tests
```bash
go test -v ./tests/integration/...
```

## Development

### Available Make Commands

```bash
make build          # Build the application
make run            # Run the application
make test           # Run tests
make test-coverage  # Run tests with coverage
make clean          # Clean build files
make docker-up      # Start Docker services
make docker-down    # Stop Docker services
make docker-build   # Build Docker image
make migrate-up     # Run database migrations
make migrate-down   # Rollback migrations
make gen-proto      # Generate gRPC code
make deps           # Install dependencies
make lint           # Lint code
make fmt            # Format code
make dev-setup      # Setup development environment
make prod-build     # Production build
```

### Code Style

- Follow Go conventions and best practices
- Use gofmt for code formatting
- Run golangci-lint for linting
- Write tests for all business logic
- Use meaningful variable and function names
- Add comments for complex logic

### Adding New Features

1. Define the entity in `internal/entity/`
2. Create DTOs in `internal/dto/`
3. Implement repository interface and implementation
4. Implement service interface and implementation
5. Create HTTP handlers
6. Add routes to main.go
7. Add validation middleware
8. Write tests
9. Update documentation

## Docker

### Development with Docker

```bash
# Start all services
make docker-up

# View logs
docker-compose logs -f app

# Stop services
make docker-down
```

### Production Deployment

```bash
# Build production image
make prod-build

# Build Docker image
make docker-build

# Deploy with Docker Compose
docker-compose -f docker-compose.prod.yml up -d
```

## Architecture Principles

### Clean Architecture

The project follows clean architecture principles:

- **Entities**: Core business objects
- **Use Cases**: Application business rules (Services)
- **Interface Adapters**: Controllers, gateways, presenters (Handlers, Repositories)
- **Frameworks & Drivers**: Web, database, external interfaces

### SOLID Principles

- **Single Responsibility**: Each struct/interface has one reason to change
- **Open/Closed**: Open for extension, closed for modification
- **Liskov Substitution**: Subtypes must be substitutable for their base types
- **Interface Segregation**: Many client-specific interfaces rather than one general-purpose interface
- **Dependency Inversion**: Depend on abstractions, not concretions

### Design Patterns

- **Repository Pattern**: Data access abstraction
- **Dependency Injection**: Loose coupling between components
- **Factory Pattern**: Object creation
- **Middleware Pattern**: Cross-cutting concerns
- **Observer Pattern**: Event handling (logging, caching)

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support and questions:
- Create an issue in the repository
- Check the documentation
- Review the test files for usage examples

## Changelog

### v1.0.0
- Initial release
- Basic CRUD operations
- Authentication middleware
- Database integration
- Redis caching
- gRPC support
- Docker support
- Comprehensive testing
