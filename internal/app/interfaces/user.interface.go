package interfaces

import (
	"context"

	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

type IUserRepository interface {
	GetUser(ctx context.Context, int64 int64) (*entity.User, error)
}
