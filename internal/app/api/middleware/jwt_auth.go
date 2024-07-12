package middleware

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/app/api/common"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
	"github.com/voltgizerz/POS-restaurant/internal/constants"
)

const (
	tokenTypeJWT   = "Bearer"
	headerAuthName = "Authorization"
)

type JWTAuth struct {
	AuthService ports.IJWTAuth
}

func NewJWTAuthMiddleware(authService ports.IJWTAuth) JWTAuth {
	return JWTAuth{
		AuthService: authService,
	}
}

func (m *JWTAuth) AuthorizeAccess() fiber.Handler {
	span, _ := opentracing.StartSpanFromContext(context.Background(), "middleware.AuthorizeAccess")
	defer span.Finish()

	return func(c fiber.Ctx) error {
		authHeader := c.Get(headerAuthName)
		if authHeader == "" {
			return common.SendErrorResp(c, fiber.StatusUnauthorized, "Missing Authorization header")
		}

		// Check if the token type is Bearer
		tokenType, tokenValue, err := parseAuthHeader(authHeader)
		if err != nil || tokenType != tokenTypeJWT {
			return common.SendErrorResp(c, fiber.StatusUnauthorized, "Invalid authorization header format")
		}

		// Verify JWT token using AuthService.VerifyToken
		_, claims, err := m.AuthService.VerifyToken(c.UserContext(), tokenValue)
		if err != nil {
			return common.SendErrorResp(c, fiber.StatusUnauthorized, "Invalid token")
		}

		requestID := uuid.New().String()

		// Store the UUID in the context for access in handlers
		c.Locals(constants.CTXKeyRequestID, requestID)

		userID := claims["user_id"].(float64)
		username := claims["username"].(string)
		roleID := claims["role_id"].(float64)
		isActive := claims["is_active"].(bool)

		// Create a new Go context with user information
		ctx := context.WithValue(c.UserContext(), constants.CTXKeyUserID, userID)
		ctx = context.WithValue(ctx, constants.CTXKeyUsername, username)
		ctx = context.WithValue(ctx, constants.CTXKeyRoleID, roleID)
		ctx = context.WithValue(ctx, constants.CTXKeyIsActive, isActive)
		ctx = context.WithValue(ctx, constants.CTXKeyRequestID, requestID)

		// Replace Fiber's context with the new context containing user information
		c.SetUserContext(ctx)

		return c.Next()
	}
}

// Function to parse Authorization header
func parseAuthHeader(authHeader string) (string, string, error) {
	if authHeader == "" {
		return "", "", errors.New("authorization header is empty")
	}

	splitToken := authHeader
	if len(splitToken) < 7 {
		return "", "", errors.New("authorization header format must be Bearer <token>")
	}

	return splitToken[:6], splitToken[7:], nil
}
