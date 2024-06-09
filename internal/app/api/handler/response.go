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

// sendErrorResp generates and sends error response
func sendErrorResponse(c fiber.Ctx, statusCode int, errorMessage string) error {
	response := errorResponse{
		Code:    statusCode,
		Message: errorMessage,
	}

	return c.Status(statusCode).JSON(response)
}

// sendSuccessResp generates and sends success response with dynamic data
func sendSuccessResponse(c fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := successResponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}

	return c.Status(statusCode).JSON(response)
}
