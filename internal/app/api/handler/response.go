package handler

import "github.com/gofiber/fiber/v3"

type (
	errorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	successResponse struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

// SendErrorResp generates and sends error response
func SendErrorResp(c fiber.Ctx, statusCode int, errorMessage string) error {
	response := errorResponse{
		Code:    statusCode,
		Message: errorMessage,
	}

	return c.Status(statusCode).JSON(response)
}

// SendSuccessResp generates and sends success response with dynamic data
func SendSuccessResp(c fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := successResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}

	return c.Status(statusCode).JSON(response)
}
