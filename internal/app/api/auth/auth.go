package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

const (
	authType = "Bearer"
)

type AuthJWT struct {
	SecretKey            string
	ExpireDurationInHour int
}

func NewAuthJWT(secretKey string) ports.IAuth {
	return &AuthJWT{
		SecretKey:            secretKey,
		ExpireDurationInHour: 24,
	}
}

func (a *AuthJWT) CreateToken(user *entity.User) (*entity.CreateTokenResponse, error) {
	expiredAt := time.Now().Add(time.Hour * time.Duration(a.ExpireDurationInHour))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":         user.ID,
			"expired_at": expiredAt.Unix(),
		})

	tokenString, err := token.SignedString(a.SecretKey)
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

func (a *AuthJWT) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return a.SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
