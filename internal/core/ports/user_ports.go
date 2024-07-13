package ports

import (
	"context"
	"database/sql"

	"github.com/voltgizerz/POS-restaurant/internal/core/models"
)

//go:generate mockgen -source=./internal/adapters/ports/user_ports.go -destination=./internal/mocks/mocks_user.go -package=mocks
type IUserRepository interface {
	GetUserByUsernameAndPassword(ctx context.Context, username string, hashPassword string) (*models.UserORM, error)
	GetUserByUsername(ctx context.Context, username string) (*models.UserORM, error)
	RegisterUser(ctx context.Context, tx *sql.Tx, userData models.UserORM) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*models.UserORM, error)
}
