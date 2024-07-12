package ports

import (
	"context"

	"github.com/voltgizerz/POS-restaurant/internal/core/models"
)

//go:generate mockgen -source=./internal/adapters/ports/user_ports.go -destination=./internal/mocks/mocks_user.go -package=mocks
type IUserRepository interface {
	GetUserByUsernameAndPassword(ctx context.Context, username string, hashPassword string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	RegisterUser(ctx context.Context, userData models.User) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
