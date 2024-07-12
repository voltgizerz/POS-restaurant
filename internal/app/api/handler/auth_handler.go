package handler

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/opentracing/opentracing-go"

	"github.com/voltgizerz/POS-restaurant/internal/app/api/common"
	"github.com/voltgizerz/POS-restaurant/internal/constants"
	"github.com/voltgizerz/POS-restaurant/internal/core/entity"
	"github.com/voltgizerz/POS-restaurant/internal/core/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/core/ports"
	"github.com/voltgizerz/POS-restaurant/internal/utils"
)

type AuthHandler struct {
	authService ports.IAuthService
}

func NewAuthHandler(i interactor.AuthHandler) *AuthHandler {
	return &AuthHandler{
		authService: i.AuthService,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.UserContext(), "handler.AuthHandler.Login")
	defer span.Finish()

	req := &entity.LoginRequest{}

	err := c.Bind().Body(req)
	if err != nil {
		return common.SendErrorResp(c, fiber.StatusBadRequest, constants.ErrMsgInvalidUsernameAndPassword)
	}

	err = validator.New().StructCtx(ctx, req)
	if err != nil {
		err = utils.GetFirstValidatorError(err)

		return common.SendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	dataLogin, err := h.authService.Login(ctx, req.Username, req.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return common.SendErrorResp(c, fiber.StatusUnauthorized, constants.ErrMsgUsernameNotFound)
		}

		return common.SendErrorResp(c, fiber.StatusUnauthorized, constants.ErrMsgInvalidUsernameOrPassword)
	}

	return common.SendSuccessResp(c, fiber.StatusOK, "Success", dataLogin)
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.UserContext(), "handler.AuthHandler.Register")
	defer span.Finish()

	req := &entity.RegisterRequest{}

	err := c.Bind().Body(req)
	if err != nil {
		return common.SendErrorResp(c, fiber.StatusBadRequest, "Invalid request body.")
	}

	err = validator.New().StructCtx(ctx, req)
	if err != nil {
		err = utils.GetFirstValidatorError(err)

		return common.SendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	if req.Password != req.ConfirmPassword {
		return common.SendErrorResp(c, fiber.StatusBadRequest, "Password mismatch")
	}

	userData := &entity.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Username: req.Username,
	}

	result, err := h.authService.Register(ctx, *userData)
	if err != nil {
		return common.SendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	res := map[string]int64{
		"user_id": result,
	}

	return common.SendSuccessResp(c, fiber.StatusCreated, "Account created succesfully.", res)
}
