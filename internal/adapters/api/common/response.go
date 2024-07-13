package common

import (
	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/internal/constants"
)

type (
	ErrorResponse struct {
		Success   bool   `json:"success"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	}

	SuccessResponse struct {
		Success   bool        `json:"success"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
		RequestID string      `json:"request_id"`
	}
)

// WriteErrorJSON generates and sends error response
func WriteErrorJSON(c fiber.Ctx, statusCode int, errorMessage string) error {
	response := ErrorResponse{
		Success:   false,
		Message:   errorMessage,
		RequestID: c.Locals(constants.CTXKeyRequestID).(string),
	}

	return c.Status(statusCode).JSON(response)
}

// WriteSuccessJSON generates and sends success response with dynamic data
func WriteSuccessJSON(c fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := SuccessResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		RequestID: c.Locals(constants.CTXKeyRequestID).(string),
	}

	return c.Status(statusCode).JSON(response)
}
