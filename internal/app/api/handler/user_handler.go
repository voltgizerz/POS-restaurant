package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

type UserHandler struct {
	authService ports.IAuth
	userService ports.IUserService
}

func NewUserHandler(i interactor.UserHandler) *UserHandler {
	return &UserHandler{
		authService: i.Auth,
		userService: i.UserService,
	}
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	// TODO

	//  Hit db check user and password

	// success ?
	// h.userService

	return nil
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	// TODO
	return nil
}
