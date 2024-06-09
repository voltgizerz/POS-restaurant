package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/app/constants"
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
		return sendErrorResp(c, fiber.StatusBadRequest, constants.ErrMsgInvalidUsernameAndPassword)
	}

	if req.Username == "" || req.Password == "" {
		return sendErrorResp(c, fiber.StatusBadRequest, constants.ErrMsgUsernameOrPasswordRequired)
	}

	dataUser, err := h.userService.Login(ctx, req.Username, req.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return sendErrorResp(c, fiber.StatusUnauthorized, constants.ErrMsgUsernameNotFound)
		}

		return sendErrorResp(c, fiber.StatusUnauthorized, constants.ErrMsgInvalidUsernameOrPassword)
	}

	token, err := h.authService.CreateToken(dataUser)
	if err != nil {
		return sendErrorResp(c, fiber.StatusInternalServerError, constants.ErrMsgInternalServerError)
	}

	return sendSuccessResp(c, fiber.StatusOK, "Success", token)
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	span, _ := opentracing.StartSpanFromContext(c.Context(), "handler.UserHandler.Register")
	defer span.Finish()

	// TODO
	return nil
}
