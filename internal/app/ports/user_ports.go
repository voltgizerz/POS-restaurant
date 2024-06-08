package ports

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

type IUserRepository interface {
	GetUser(ctx context.Context, int64 int64) (*entity.User, error)
}

type IUserService interface {
}

type IUserHandler interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
}
