package ports

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

type IJWTAuth interface {
	CreateToken(ctx context.Context, user entity.UserORM) (*entity.CreateTokenResponse, error)
	VerifyToken(ctx context.Context, tokenString string) (*jwt.Token, jwt.MapClaims, error)
}

type IAuthHandler interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
}
