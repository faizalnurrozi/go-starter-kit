package middleware

import (
	"reflect"
	"strings"

	dto "github.com/faizalnurrozi/go-starter-kit/internal/dto/request"

	"github.com/faizalnurrozi/go-starter-kit/internal/errors"
	"github.com/faizalnurrozi/go-starter-kit/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateRequest(requestType interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create new instance of request type
		reqType := reflect.TypeOf(requestType)
		if reqType.Kind() == reflect.Ptr {
			reqType = reqType.Elem()
		}
		req := reflect.New(reqType).Interface()

		// Parse body
		if err := c.BodyParser(req); err != nil {
			return utils.SendError(c, errors.NewValidationError("Invalid request body"))
		}

		// Validate
		if err := validate.Struct(req); err != nil {
			var errorMessages []string
			for _, err := range err.(validator.ValidationErrors) {
				errorMessages = append(errorMessages, err.Field()+" "+err.Tag())
			}
			return utils.SendError(c, errors.NewValidationError("Validation failed: "+strings.Join(errorMessages, ", ")))
		}

		c.Locals("validatedRequest", req)
		return c.Next()
	}
}

func ValidateParams() fiber.Handler {
	return func(c *fiber.Ctx) error {
		params := &dto.GetUserParams{}

		if err := c.ParamsParser(params); err != nil {
			return utils.SendError(c, errors.NewValidationError("Invalid parameters"))
		}

		if err := validate.Struct(params); err != nil {
			return utils.SendError(c, errors.NewValidationError("Invalid parameters"))
		}

		c.Locals("validatedParams", params)
		return c.Next()
	}
}
