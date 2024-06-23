package middleware

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/app/api/handler"
	"github.com/voltgizerz/POS-restaurant/internal/app/constants"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

const (
	tokeTypeJWT    = "Bearer"
	headerAuthName = "Authorization"
)

type JWTAuth struct {
	AuthService ports.IAuth
}

func NewJWTAuthMiddleware(authService ports.IAuth) JWTAuth {
	return JWTAuth{
		AuthService: authService,
	}
}

func (m *JWTAuth) AuthorizeAccess() fiber.Handler {
	span, _ := opentracing.StartSpanFromContext(context.Background(), "middleware.AuthorizeAccess")
	defer span.Finish()

	return func(ctx fiber.Ctx) error {
		authHeader := ctx.Get(headerAuthName)
		if authHeader == "" {
			return handler.SendErrorResp(ctx, fiber.StatusUnauthorized, "Missing Authorization header")
		}

		// Check if the token type is Bearer
		tokenType, tokenValue, err := parseAuthHeader(authHeader)
		if err != nil || tokenType != tokeTypeJWT {
			return handler.SendErrorResp(ctx, fiber.StatusUnauthorized, "Invalid authorization header format")
		}

		// Verify JWT token using AuthService.VerifyToken
		_, claims, err := m.AuthService.VerifyToken(ctx.Context(), tokenValue)
		if err != nil {
			return handler.SendErrorResp(ctx, fiber.StatusUnauthorized, "Invalid token")
		}

		// Extract user ID from claims
		userID := claims["user_id"].(float64)
		username := claims["username"].(string)
		roleID := claims["role_id"].(float64)
		isActive := claims["is_active"].(bool)

		// Set user information in request context
		ctx.Locals(constants.CTXKeyUserID, int64(userID))
		ctx.Locals(constants.CTXKeyUsername, username)
		ctx.Locals(constants.CTXKeyRoleID, int64(roleID))
		ctx.Locals(constants.CTXKeyIsActive, isActive)

		return ctx.Next()
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
