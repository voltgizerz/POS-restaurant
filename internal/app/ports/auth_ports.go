package ports

import (
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

type IAuth interface {
	CreateToken(user *entity.UserORM) (*entity.CreateTokenResponse, error)
	VerifyToken(tokenString string) error
}
