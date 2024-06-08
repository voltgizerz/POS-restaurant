package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

type UserHandler struct {
	userService ports.IUserService
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	// TODO
	return nil
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	// TODO
	return nil
}
