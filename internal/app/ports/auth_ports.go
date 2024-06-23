package ports

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

type IAuth interface {
	CreateToken(ctx context.Context, user entity.UserORM) (*entity.CreateTokenResponse, error)
	VerifyToken(ctx context.Context, tokenString string) (*jwt.Token, jwt.MapClaims, error)
}
