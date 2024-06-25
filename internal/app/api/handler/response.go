package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/internal/app/constants"
)

type (
	errorResponse struct {
		Success   bool   `json:"success"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	}

	successResponse struct {
		Success   bool        `json:"success"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
		RequestID string      `json:"request_id"`
	}
)

// SendErrorResp generates and sends error response
func SendErrorResp(c fiber.Ctx, statusCode int, errorMessage string) error {
	response := errorResponse{
		Success:   false,
		Message:   errorMessage,
		RequestID: c.Locals(constants.CTXKeyRequestID).(string),
	}

	return c.Status(statusCode).JSON(response)
}

// SendSuccessResp generates and sends success response with dynamic data
func SendSuccessResp(c fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := successResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		RequestID: c.Locals(constants.CTXKeyRequestID).(string),
	}

	return c.Status(statusCode).JSON(response)
}
