package ports

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

type IUserRepository interface {
	GetUserByUsernameAndPassword(ctx context.Context, username string, hashPassword string) (*entity.UserORM, error)
	RegisterUser(ctx context.Context, username string, hashPassword string, Email string, Name string) (bool, error)
}

type IUserService interface {
	Login(ctx context.Context, username string, password string) (*entity.User, error)
	Register(ctx context.Context, username string, email string, password string, confirmPass string, name string) (bool, error)
}

type IUserHandler interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
}
