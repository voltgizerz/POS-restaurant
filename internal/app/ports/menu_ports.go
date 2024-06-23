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
	UpdateActiveMenuBatchByUserID(c fiber.Ctx) error
	UpdateActiveMenuByMenuID(c fiber.Ctx) error
}

type IMenuService interface {
	RegisterMenu(ctx context.Context, menuData entity.Menu) (int64, error)
	GetMenu(ctx context.Context, idMenu int64) ([]*entity.MenuResponse, error)
	UpdateActiveMenuID(ctx context.Context, idMenu int64) (int64, error)
	UpdateActiveMenuBatchUserID(ctx context.Context, idUser int64) (int64, error)
	UpdateMenuID(ctx context.Context, menuData *entity.Menu) (int64, error)
}

type IMenuRepository interface {
	FetchMenuById(ctx context.Context, idUser int64) ([]*entity.MenuOrm, error)
	AddMenu(ctx context.Context, menuData *entity.MenuOrm) (int64, error)
	UpdateActiveMenuByMenuID(ctx context.Context, idMenu int64) (int64, error)
	UpdateActiveMenuBatchUser(ctx context.Context, idUser int64) (int64, error)
	UpdateMenuByMenuID(ctx context.Context, menuData *entity.MenuOrm) (int64, error)
}
