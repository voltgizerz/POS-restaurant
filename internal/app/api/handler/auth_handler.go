package handler

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/opentracing/opentracing-go"

	"github.com/voltgizerz/POS-restaurant/internal/app/constants"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

type AuthHandler struct {
	userService ports.IUserService
}

func NewAuthHandler(i interactor.UserHandler) *AuthHandler {
	return &AuthHandler{
		userService: i.UserService,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.AuthHandler.Login")
	defer span.Finish()

	req := &loginRequest{}

	err := c.Bind().Body(req)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, constants.ErrMsgInvalidUsernameAndPassword)
	}

	err = validator.New().StructCtx(ctx, req)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		return SendErrorResp(c, fiber.StatusBadRequest, validationErrors.Error())
	}

	userLoginData, err := h.userService.Login(ctx, req.Username, req.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return SendErrorResp(c, fiber.StatusUnauthorized, constants.ErrMsgUsernameNotFound)
		}

		return SendErrorResp(c, fiber.StatusUnauthorized, constants.ErrMsgInvalidUsernameOrPassword)
	}

	return SendSuccessResp(c, fiber.StatusOK, "Success", userLoginData)
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.AuthHandler.Register")
	defer span.Finish()

	req := &registerRequest{}
	err := c.Bind().Body(req)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, "Invalid request body.")
	}

	err = validator.New().StructCtx(ctx, req)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		return SendErrorResp(c, fiber.StatusBadRequest, validationErrors.Error())
	}

	if req.Password != req.ConfirmPassword {
		return SendErrorResp(c, fiber.StatusBadRequest, "Password mismatch")
	}

	userData := &entity.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Username: req.Username,
	}

	result, err := h.userService.Register(ctx, *userData)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}
	res := map[string]int64{
		"user_id": result,
	}

	return SendSuccessResp(c, fiber.StatusCreated, "Account created succesfully.", res)
}
