package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string, details ...string) *AppError {
	err := &AppError{
		Code:    code,
		Message: message,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}

// Predefined errors
func NewValidationError(message string) *AppError {
	return NewAppError(http.StatusBadRequest, message)
}

func NewNotFoundError(resource string) *AppError {
	return NewAppError(http.StatusNotFound, fmt.Sprintf("%s not found", resource))
}

func NewUnauthorizedError() *AppError {
	return NewAppError(http.StatusUnauthorized, "Unauthorized")
}

func NewInternalError(message string) *AppError {
	return NewAppError(http.StatusInternalServerError, message)
}

func NewBusinessError(message string) *AppError {
	return NewAppError(404, message) // Using 404 for business errors as requested
}
