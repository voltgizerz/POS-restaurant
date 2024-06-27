package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

const (
	authType          = "Bearer"
	expireDurationJWT = 24 // Hour
)

type AuthJWT struct {
	SecretKey      string
	ExpireDuration int
}

func NewAuthJWT(secretKey string) ports.IJWTAuth {
	return &AuthJWT{
		SecretKey:      secretKey,
		ExpireDuration: expireDurationJWT,
	}
}

func (a *AuthJWT) CreateToken(ctx context.Context, user entity.UserORM) (*entity.CreateTokenResponse, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "auth.CreateToken")
	defer span.Finish()

	expiredAt := time.Now().Add(time.Hour * time.Duration(a.ExpireDuration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":    user.ID,
			"username":   user.Username,
			"role_id":    user.RoleID,
			"is_active":  user.IsActive,
			"expired_at": expiredAt.Unix(),
		})

	tokenString, err := token.SignedString([]byte(a.SecretKey))
	if err != nil {
		return nil, err
	}

	resp := &entity.CreateTokenResponse{
		Token:     tokenString,
		ExpiredAt: expiredAt,
		TokenType: authType,
	}

	return resp, nil
}

func (a *AuthJWT) VerifyToken(ctx context.Context, tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "auth.VerifyToken")
	defer span.Finish()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.SecretKey), nil
	})

	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, errors.New("invalid token claims")
	}

	return token, claims, nil
}
