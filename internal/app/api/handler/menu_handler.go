package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/app/constants"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

type MenuHandler struct {
	menuService ports.IMenuService
}

func NewMenuHandler(i interactor.MenuHandler) *MenuHandler {
	return &MenuHandler{
		menuService: i.MenuService,
	}
}

func (h *MenuHandler) AddMenu(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.MenuHandler.AddMenu")
	defer span.Finish()

	req := &addMenuRequest{}

	err := c.Bind().Body(req)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, "Error data menu")
	}

	priceConvert, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	menuData := &entity.Menu{
		Name:      req.Name,
		Thumbnail: req.Thumbnail,
		UserID:    req.UserID,
		IsActive:  req.IsActive,
		Price:     priceConvert,
	}

	result, err := h.menuService.RegisterMenu(ctx, *menuData)
	if err != nil {
		return SendErrorResp(c, fiber.StatusUnauthorized, err.Error())
	}

	responseMsg := map[string]int64{
		"menu_id": result,
	}

	return SendSuccessResp(c, fiber.StatusOK, "Success", responseMsg)
}

func (h *MenuHandler) GetMenuByUserID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.MenuHandler.GetMenuByUserID")
	defer span.Finish()

	userID := c.Params("user_id")
	convertUserIDtoInt, err := strconv.Atoi(userID)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.menuService.GetMenu(ctx, int64(convertUserIDtoInt))
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, constants.ErrMsgMenuNotFound)
	}

	return SendSuccessResp(c, fiber.StatusOK, "Success", result)
}

func (h *MenuHandler) UpdateMenuByMenuID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.MenuHandler.GetMenuByUserID")
	defer span.Finish()

	req := &updateMenuRequest{}
	err := c.Bind().Body(req)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, "Invalid request body.")
	}

	menuID := c.Params("menu_id")
	convertMenuIDtoInt, err := strconv.Atoi(menuID)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	priceConvert, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	menuData := entity.Menu{
		ID:        int64(convertMenuIDtoInt),
		Name:      req.Name,
		UserID:    req.UserID,
		Thumbnail: req.Thumbnail,
		Price:     priceConvert,
		IsActive:  req.IsActive,
	}

	result, err := h.menuService.UpdateMenuID(ctx, menuData)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, constants.ErrMsgFailedUpdateMenu)
	}

	return SendSuccessResp(c, fiber.StatusOK, "Success", result)
}

func (h *MenuHandler) UpdateActiveMenuBatchByUserID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.MenuHandler.UpdateActiveMenuBatchByUserID")
	defer span.Finish()

	userID := c.Params("user_id")
	convertUserIDtoInt, err := strconv.Atoi(userID)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.menuService.UpdateActiveMenuBatchUserID(ctx, int64(convertUserIDtoInt))
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, fmt.Sprintf(constants.ErrMsgFailedDeleteMenu, " delete batch by user id"))
	}

	return SendSuccessResp(c, fiber.StatusOK, "Success", result)
}

func (h *MenuHandler) UpdateActiveMenuByMenuID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.Context(), "handler.MenuHandler.DeleteMenuByMenuId")
	defer span.Finish()

	menuID := c.Params("menu_id")
	convertMenuIDtoInt, err := strconv.Atoi(menuID)
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.menuService.UpdateActiveMenuID(ctx, int64(convertMenuIDtoInt))
	if err != nil {
		return SendErrorResp(c, fiber.StatusBadRequest, fmt.Sprintf(constants.ErrMsgFailedDeleteMenu, " delete by menu id"))
	}

	return SendSuccessResp(c, fiber.StatusOK, "Success", result)
}
