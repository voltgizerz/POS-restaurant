package ports

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/internal/core/entity"
	"github.com/voltgizerz/POS-restaurant/internal/core/models"
)

//go:generate mockgen -source=./internal/adapters/ports/menu_ports.go -destination=./internal/mocks/mocks_menu.go -package=mocks
type IMenuHandler interface {
	AddMenu(c fiber.Ctx) error
	GetMenuByUserID(c fiber.Ctx) error
	UpdateMenuByMenuID(c fiber.Ctx) error
	UpdateActiveMenuBatchByUserID(c fiber.Ctx) error
	UpdateActiveMenuByMenuID(c fiber.Ctx) error
}

type IMenuService interface {
	RegisterMenu(ctx context.Context, menuData entity.Menu) (int64, error)
	GetMenu(ctx context.Context, menuID int64) ([]*entity.MenuResponse, error)
	UpdateActiveMenuID(ctx context.Context, menuID int64) (int64, error)
	UpdateActiveMenuBatchUserID(ctx context.Context, userID int64) (int64, error)
	UpdateMenuID(ctx context.Context, menuData entity.Menu) (int64, error)
}

type IMenuRepository interface {
	FetchMenuById(ctx context.Context, userID int64) ([]*models.MenuORM, error)
	AddMenu(ctx context.Context, menuData models.MenuORM) (int64, error)
	UpdateActiveMenuByMenuID(ctx context.Context, tx *sql.Tx, menuID int64) (int64, error)
	UpdateActiveMenuBatchUser(ctx context.Context, tx *sql.Tx, userID int64) (int64, error)
	UpdateMenuByMenuID(ctx context.Context, tx *sql.Tx, menuData *models.MenuORM) (int64, error)
}
