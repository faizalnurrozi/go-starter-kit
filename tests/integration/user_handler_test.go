package integration

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	dto "github.com/faizalnurrozi/go-starter-kit/internal/dto/request"
	"github.com/faizalnurrozi/go-starter-kit/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Create_Integration(t *testing.T) {
	// Setup test database, repositories, services, etc.
	// This is a simplified example

	app := fiber.New()
	userHandler := handler.NewUserHandler(nil) // Replace with actual service

	app.Post("/users", userHandler.Create)

	req := dto.CreateUserRequest{
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
