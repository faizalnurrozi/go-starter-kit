package handler

import (
	"strconv"

	dto "github.com/faizalnurrozi/go-starter-kit/internal/dto/request"
	"github.com/faizalnurrozi/go-starter-kit/internal/service/interfaces"
	"github.com/faizalnurrozi/go-starter-kit/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	req := c.Locals("validatedRequest").(*dto.CreateUserRequest)

	user, err := h.userService.Create(c.Context(), req)
	if err != nil {
		return utils.SendError(c, err)
	}

	return utils.SendSuccess(c, user)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	params := c.Locals("validatedParams").(*dto.GetUserParams)

	user, err := h.userService.GetByID(c.Context(), params.ID)
	if err != nil {
		return utils.SendError(c, err)
	}

	return utils.SendSuccess(c, user)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	users, err := h.userService.GetAll(c.Context(), limit, offset)
	if err != nil {
		return utils.SendError(c, err)
	}

	return utils.SendSuccess(c, users)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	params := c.Locals("validatedParams").(*dto.GetUserParams)
	req := c.Locals("validatedRequest").(*dto.UpdateUserRequest)

	user, err := h.userService.Update(c.Context(), params.ID, req)
	if err != nil {
		return utils.SendError(c, err)
	}

	return utils.SendSuccess(c, user)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	params := c.Locals("validatedParams").(*dto.GetUserParams)

	err := h.userService.Delete(c.Context(), params.ID)
	if err != nil {
		return utils.SendError(c, err)
	}

	return utils.SendSuccess(c, map[string]string{"message": "User deleted successfully"})
}
