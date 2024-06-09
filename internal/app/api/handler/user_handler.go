package handler

import (
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/app/constants"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

type UserHandler struct {
	authService ports.IAuth
	userService ports.IUserService
}

func NewUserHandler(i interactor.UserHandler) *UserHandler {
	return &UserHandler{
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

	userLoginData, err := h.userService.Login(ctx, req.Username, req.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return sendErrorResp(c, fiber.StatusUnauthorized, constants.ErrMsgUsernameNotFound)
		}

		return sendErrorResp(c, fiber.StatusUnauthorized, constants.ErrMsgInvalidUsernameOrPassword)
	}

	return sendSuccessResp(c, fiber.StatusOK, "Success", userLoginData)
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.UserHandler.Register")
	defer span.Finish()

	req := &registerRequest{}
	err := c.Bind().Body(req)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, "Invalid request body.")
	}
	if req.Password != req.ConfirmPassword {
		return sendErrorResp(c, fiber.StatusBadRequest, "Password Mismatch")
	}

	userData := &entity.UserORM{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Username: req.Username,
	}

	result, err := h.userService.Register(ctx, *userData)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}
	return sendSuccessResp(c, fiber.StatusOK, "Account Added Succesfully.", result)
}
