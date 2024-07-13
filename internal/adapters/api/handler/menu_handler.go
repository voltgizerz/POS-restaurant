package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/adapters/api/common"
	"github.com/voltgizerz/POS-restaurant/internal/constants"
	"github.com/voltgizerz/POS-restaurant/internal/core/entity"
	"github.com/voltgizerz/POS-restaurant/internal/core/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/core/ports"
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
	span, ctx := opentracing.StartSpanFromContext(c.UserContext(), "handler.MenuHandler.AddMenu")
	defer span.Finish()

	req := &addMenuRequest{}

	err := c.Bind().Body(req)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, "Error data menu")
	}

	priceConvert, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, err.Error())
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
		return common.WriteErrorJSON(c, fiber.StatusUnauthorized, err.Error())
	}

	responseMsg := map[string]int64{
		"menu_id": result,
	}

	return common.WriteSuccessJSON(c, fiber.StatusOK, "Success", responseMsg)
}

func (h *MenuHandler) GetMenuByUserID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.UserContext(), "handler.MenuHandler.GetMenuByUserID")
	defer span.Finish()

	userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.menuService.GetMenu(ctx, userID)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, constants.ErrMsgMenuNotFound)
	}

	return common.WriteSuccessJSON(c, fiber.StatusOK, "Success", result)
}

func (h *MenuHandler) UpdateMenuByMenuID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.UserContext(), "handler.MenuHandler.GetMenuByUserID")
	defer span.Finish()

	req := &updateMenuRequest{}
	err := c.Bind().Body(req)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, "Invalid request body.")
	}

	menuID, err := strconv.ParseInt(c.Params("menu_id"), 10, 64)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, err.Error())
	}

	priceConvert, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, err.Error())
	}

	menuData := entity.Menu{
		ID:        menuID,
		Name:      req.Name,
		UserID:    req.UserID,
		Thumbnail: req.Thumbnail,
		Price:     priceConvert,
		IsActive:  req.IsActive,
	}

	result, err := h.menuService.UpdateMenuID(ctx, menuData)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, constants.ErrMsgFailedUpdateMenu)
	}

	return common.WriteSuccessJSON(c, fiber.StatusOK, "Success", result)
}

func (h *MenuHandler) UpdateActiveMenuBatchByUserID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.UserContext(), "handler.MenuHandler.UpdateActiveMenuBatchByUserID")
	defer span.Finish()

	userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.menuService.UpdateActiveMenuBatchUserID(ctx, userID)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, fmt.Sprintf(constants.ErrMsgFailedDeleteMenu, " delete batch by user id"))
	}

	return common.WriteSuccessJSON(c, fiber.StatusOK, "Success", result)
}

func (h *MenuHandler) UpdateActiveMenuByMenuID(c fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(c.UserContext(), "handler.MenuHandler.DeleteMenuByMenuId")
	defer span.Finish()

	menuID, err := strconv.ParseInt(c.Params("menu_id"), 10, 64)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := h.menuService.UpdateActiveMenuID(ctx, menuID)
	if err != nil {
		return common.WriteErrorJSON(c, fiber.StatusBadRequest, fmt.Sprintf(constants.ErrMsgFailedDeleteMenu, "delete by menu id"))
	}

	return common.WriteSuccessJSON(c, fiber.StatusOK, "Success", result)
}
