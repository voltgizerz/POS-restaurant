package ports

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

//go:generate mockgen -source=./internal/app/ports/auth_ports.go -destination=./internal/mocks/mocks_auth.go -package=mocks
type IJWTAuth interface {
	CreateToken(ctx context.Context, user entity.UserORM) (*entity.CreateTokenResponse, error)
	VerifyToken(ctx context.Context, tokenString string) (*jwt.Token, jwt.MapClaims, error)
}

type IAuthHandler interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
}

type IAuthService interface {
	Login(ctx context.Context, username string, password string) (*entity.LoginResponse, error)
	Register(ctx context.Context, userData entity.User) (int64, error)
}
