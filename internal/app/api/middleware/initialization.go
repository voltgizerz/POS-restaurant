package middleware

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/voltgizerz/POS-restaurant/internal/constants"
)

// All request go here first
func Initialization(c fiber.Ctx) error {
	requestID := uuid.New().String()

	// Store the UUID in the context for access in handlers
	c.Locals(constants.CTXKeyRequestID, requestID)

	// Create a new Go context with the request ID
	ctx := context.WithValue(c.UserContext(), constants.CTXKeyRequestID, requestID)

	// Replace Fiber's context with the new context containing the UUID
	c.SetUserContext(ctx)

	return c.Next()
}
