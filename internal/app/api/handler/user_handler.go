package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/opentracing/opentracing-go"
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
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.UserHandler.Login")
	defer span.Finish()

	req := &loginRequest{}

	err := c.Bind().Body(req)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Invalid request body. Please provide both username and password.")
	}

	// Check if the username and password are provided
	if req.Username == "" || req.Password == "" {
		return sendErrorResponse(c, fiber.StatusBadRequest, "Username and password are required.")
	}

	dataUser, err := h.userService.Login(ctx, req.Username, req.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return sendErrorResponse(c, fiber.StatusUnauthorized, "Username not found.")
		}

		return sendErrorResponse(c, fiber.StatusUnauthorized, "Invalid username or password.")
	}

	token, err := h.authService.CreateToken(dataUser)
	if err != nil {
		return sendErrorResponse(c, fiber.StatusInternalServerError, "Failed create user token.")
	}

	return sendSuccessResponse(c, fiber.StatusOK, "Login successful.", token)
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	span, _ := opentracing.StartSpanFromContext(c.Context(), "handler.UserHandler.Register")
	defer span.Finish()

	// TODO
	return nil
}
