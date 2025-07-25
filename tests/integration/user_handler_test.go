package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	req "github.com/faizalnurrozi/go-starter-kit/internal/dto/request"
	res "github.com/faizalnurrozi/go-starter-kit/internal/dto/response"
	"github.com/faizalnurrozi/go-starter-kit/internal/handler"
	"github.com/faizalnurrozi/go-starter-kit/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type dummyUserService struct{}

// Create implements interfaces.UserService.
func (s *dummyUserService) Create(ctx context.Context, req *req.CreateUserRequest) (*res.UserResponse, error) {
	return &res.UserResponse{
		ID:    1,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

// Delete implements interfaces.UserService.
func (s *dummyUserService) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// GetAll implements interfaces.UserService.
func (s *dummyUserService) GetAll(ctx context.Context, limit int, offset int) ([]*res.UserResponse, error) {
	panic("unimplemented")
}

// GetByID implements interfaces.UserService.
func (s *dummyUserService) GetByID(ctx context.Context, id uint) (*res.UserResponse, error) {
	panic("unimplemented")
}

// Update implements interfaces.UserService.
func (s *dummyUserService) Update(ctx context.Context, id uint, req *req.UpdateUserRequest) (*res.UserResponse, error) {
	panic("unimplemented")
}

func (s *dummyUserService) CreateUser(input req.CreateUserRequest) (interface{}, error) {
	return fiber.Map{"message": "user created successfully"}, nil
}

func TestUserHandler_Create_Integration(t *testing.T) {
	// Setup test database, repositories, services, etc.
	// This is a simplified example

	app := fiber.New()
	userService := &dummyUserService{}
	userHandler := handler.NewUserHandler(userService) // Replace with actual service

	app.Post("/users", middleware.ValidateRequest(&req.CreateUserRequest{}), userHandler.Create)

	req := req.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(httpReq)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
