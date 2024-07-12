package ports

import (
	"context"

	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

//go:generate mockgen -source=./internal/app/ports/user_ports.go -destination=./internal/mocks/mocks_user.go -package=mocks
type IUserRepository interface {
	GetUserByUsernameAndPassword(ctx context.Context, username string, hashPassword string) (*entity.UserORM, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.UserORM, error)
	RegisterUser(ctx context.Context, userData entity.UserORM) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.UserORM, error)
}
