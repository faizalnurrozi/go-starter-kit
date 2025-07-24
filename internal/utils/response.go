package utils

import (
	"encoding/xml"

	"github.com/faizalnurrozi/go-starter-kit/internal/errors"
	"github.com/faizalnurrozi/go-starter-kit/internal/logger"

	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Status  string      `json:"status" xml:"status"`
	Code    int         `json:"code" xml:"code"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Error   interface{} `json:"error,omitempty" xml:"error,omitempty"`
}

func SendSuccess(c *fiber.Ctx, data interface{}) error {
	response := BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "Request successful",
		Data:    data,
	}
	return sendResponse(c, 200, response)
}

func SendError(c *fiber.Ctx, err error) error {
	var response BaseResponse

	switch e := err.(type) {
	case *errors.AppError:
		response = BaseResponse{
			Status:  "error",
			Code:    e.Code,
			Message: e.Message,
			Error:   e.Details,
		}
		return sendResponse(c, getHTTPStatus(e.Code), response)
	default:
		logger.Error("Unexpected error: ", err)
		response = BaseResponse{
			Status:  "error",
			Code:    500,
			Message: "Internal server error",
		}
		return sendResponse(c, 500, response)
	}
}

func sendResponse(c *fiber.Ctx, statusCode int, response BaseResponse) error {
	acceptHeader := c.Get("Accept")

	switch {
	case fiber.MIMEApplicationXML == acceptHeader:
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationXML)
		xmlData, _ := xml.Marshal(response)
		return c.Status(statusCode).Send(xmlData)
	default:
		return c.Status(statusCode).JSON(response)
	}
}

func getHTTPStatus(code int) int {
	switch code {
	case 404: // Business error
		return 422 // Unprocessable Entity for business logic errors
	case 401:
		return 401
	case 400:
		return 400
	default:
		return 500
	}
}
