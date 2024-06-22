package ports

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
)

type IMenuHandler interface {
	AddMenu(c fiber.Ctx) error
	GetMenuByUserID(c fiber.Ctx) error
	UpdateMenuByMenuID(c fiber.Ctx) error
	DeleteMenuBatchByUserID(c fiber.Ctx) error
	DeleteMenuByMenuId(c fiber.Ctx) error
}

type IMenuService interface {
	RegisterMenu(ctx context.Context, menuData entity.Menu) (int64, error)
	GetMenu(ctx context.Context, idMenu int64) ([]*entity.MenuResponse, error)
	DeleteMenuID(ctx context.Context, idMenu int64) (int64, error)
	DeleteMenuBatchUserID(ctx context.Context, idUser int64) (int64, error)
	UpdateMenuID(ctx context.Context, menuData *entity.Menu) (int64, error)
}

type IMenuRepository interface {
	FetchMenuById(ctx context.Context, idUser int64) ([]*entity.MenuOrm, error)
	AddMenu(ctx context.Context, menuData *entity.MenuOrm) (int64, error)
	DeleteMenuByMenuID(ctx context.Context, idMenu int64) (int64, error)
	DeleteMenuBatchUser(ctx context.Context, idUser int64) (int64, error)
	UpdateMenuByMenuID(ctx context.Context, menuData *entity.MenuOrm) (int64, error)
}
