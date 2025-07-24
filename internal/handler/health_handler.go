package handler

import (
	"github.com/faizalnurrozi/go-starter-kit/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	data := map[string]interface{}{
		"status":  "healthy",
		"service": "github.com/faizalnurrozi/go-starter-kit",
		"version": "1.0.0",
	}

	return utils.SendSuccess(c, data)
}
