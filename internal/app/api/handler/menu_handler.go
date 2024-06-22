package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/opentracing/opentracing-go"
	"github.com/shopspring/decimal"
	"github.com/voltgizerz/POS-restaurant/internal/app/constants"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

type MenuHandler struct {
	menuService ports.IMenuService
	authService ports.IAuth
}

func NewMenuHandler(i interactor.MenuHandler) *MenuHandler {
	return &MenuHandler{
		menuService: i.MenuService,
		authService: i.AuthService,
	}
}

func (h *MenuHandler) AddMenu(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.MenuHandler.AddMenu")
	defer span.Finish()

	req := &addMenuRequest{}

	err := c.Bind().Body(req)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, "Error data menu")
	}

	convertId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	priceConvert, err := decimal.NewFromString(req.Price)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	menuData := &entity.Menu{
		Name:      req.Name,
		Thumbnail: req.Thumbnail,
		UserId:    int64(convertId),
		Price:     priceConvert,
	}

	result, err := h.menuService.RegisterMenu(ctx, *menuData)
	if err != nil {
		return sendErrorResp(c, fiber.StatusUnauthorized, err.Error())
	}

	responseMsg := map[string]int64{
		"menu_id": result,
	}

	return sendSuccessResp(c, fiber.StatusOK, "Success", responseMsg)
}

func (h *MenuHandler) GetMenuByUserID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.MenuHandler.GetMenuByUserID")
	defer span.Finish()

	req := &getMenuRequest{}
	err := c.Bind().Body(req)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, "Invalid request body.")
	}

	convertId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, "Error Convert ID")
	}

	result, err := h.menuService.GetMenu(ctx, int64(convertId))
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, constants.ErrMsgMenuNotFound)
	}

	return sendSuccessResp(c, fiber.StatusOK, "Success", result)
}

func (h *MenuHandler) UpdateMenuByMenuID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.MenuHandler.GetMenuByUserID")
	defer span.Finish()

	req := &updateMenuRequest{}
	err := c.Bind().Body(req)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, "Invalid request body.")
	}

	convertMenuid, err := strconv.Atoi(req.ID)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	convertUserid, err := strconv.Atoi(req.UserId)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	priceConvert, err := decimal.NewFromString(req.Price)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	activeConvert, err := strconv.ParseBool(req.IsActive)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	menuData := entity.Menu{
		ID:        int64(convertMenuid),
		Name:      req.Name,
		UserId:    int64(convertUserid),
		Thumbnail: req.Thumbnail,
		Price:     priceConvert,
		IsActive:  activeConvert,
	}

	result, err := h.menuService.UpdateMenuID(ctx, &menuData)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, constants.ErrMsgFailedUpdateMenu)
	}

	return sendSuccessResp(c, fiber.StatusOK, "Success", result)
}

func (h *MenuHandler) DeleteMenuBatchByUserID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.MenuHandler.DeleteMenuBatchByUserID")
	defer span.Finish()

	req := &getMenuRequest{}
	err := c.Bind().Body(req)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, "Invalid request body.")
	}

	convertId, err := strconv.Atoi(req.UserId)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, "Error Convert ID")
	}

	result, err := h.menuService.DeleteMenuBatchUserID(ctx, int64(convertId))
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, fmt.Sprintf(constants.ErrMsgFailedDeleteMenu, " Delete Batch By User Id"))
	}

	return sendSuccessResp(c, fiber.StatusOK, "Success", result)
}

func (h *MenuHandler) DeleteMenuByMenuId(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.MenuHandler.DeleteMenuByMenuId")
	defer span.Finish()

	req := &menuIdRequest{}
	err := c.Bind().Body(req)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, "Invalid request body.")
	}

	convertId, err := strconv.Atoi(req.MenuId)
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, "Error Convert ID")
	}

	result, err := h.menuService.DeleteMenuID(ctx, int64(convertId))
	if err != nil {
		return sendErrorResp(c, fiber.StatusBadRequest, fmt.Sprintf(constants.ErrMsgFailedDeleteMenu, " Delete By Menu ID"))
	}

	return sendSuccessResp(c, fiber.StatusOK, "Success", result)
}
