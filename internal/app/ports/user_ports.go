package ports

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

type IUserRepository interface {
	GetUserByUsernameAndPassword(ctx context.Context, username string, hashPassword string) (*entity.UserORM, error)
	RegisterUser(ctx context.Context, userData entity.UserORM) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (entity.UserORM, error)
}

type IUserService interface {
	Login(ctx context.Context, username string, password string) (*entity.LoginResponse, error)
	Register(ctx context.Context, userData entity.User) (int64, error)
}

type IUserHandler interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
}
