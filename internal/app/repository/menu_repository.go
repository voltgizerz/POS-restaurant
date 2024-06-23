package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/app/entity"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

const (
	queryInsertMenu                    = `INSERT INTO food_menus (name,thumbnail,price,is_active,user_id) values (?,?,?,?,?)`
	queryGetMenuByUserId               = `SELECT id,name,price,thumbnail,is_active from food_menus where user_id = ? and is_active = 1`
	queryUpdateActiveBatchMenuByUserId = `UPDATE food_menus set is_active = 0 , deleted_at = ? where user_id = ?`
	queryUpdateMenuActiveByMenuId      = `UPDATE food_menus set is_active = 0 , deleted_at = ? where id = ?`
	queryUpdateMenuByMenuId            = `UPDATE food_menus set name = ? ,price = ? , thumbnail = ? ,is_active = ? , user_id = ? , updated_at = ? where id = ?`
)

type MenuRepository struct {
	MasterDB *sqlx.DB
}

func NewMenuRepository(opts RepositoryOpts) ports.IMenuRepository {
	return &MenuRepository{
		MasterDB: opts.Database.MasterDB,
	}
}

// DeleteMenuBatchUser implements ports.IMenuRepository.
func (m *MenuRepository) UpdateActiveMenuBatchUser(ctx context.Context, idUser int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.MenuRepository.UpdateActiveMenuBatchUser")
	defer span.Finish()

	result, err := m.MasterDB.ExecContext(ctx, queryUpdateActiveBatchMenuByUserId, time.Now(), idUser)
	if err != nil {
		return 0, err
	}

	rowChanged, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowChanged <= 0 {
		return 0, errors.New("No Data Update")
	}

	return 1, nil
}

// UpdateMenuByMenuID implements ports.IMenuRepository.
func (m *MenuRepository) UpdateActiveMenuByMenuID(ctx context.Context, menuID int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.MenuRepository.UpdateActiveMenuByMenuID")
	defer span.Finish()

	result, err := m.MasterDB.ExecContext(ctx, queryUpdateMenuActiveByMenuId, time.Now(), menuID)
	if err != nil {
		return 0, err
	}

	rowChanged, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowChanged <= 0 {
		return 0, errors.New("No Data Update")
	}

	return 1, nil
}

// UpdateMenuByMenuID implements ports.IMenuRepository.
func (m *MenuRepository) UpdateMenuByMenuID(ctx context.Context, menuData *entity.MenuORM) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.MenuRepository.UpdateMenuByMenuID")
	defer span.Finish()

	result, err := m.MasterDB.ExecContext(ctx, queryUpdateMenuByMenuId, menuData.Name, menuData.Price, menuData.Thumbnail, menuData.IsActive, menuData.UserId, time.Now(), menuData.ID)
	if err != nil {
		return 0, err
	}

	rowChanged, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowChanged <= 0 {
		return 0, errors.New("No Data Update")
	}

	return 1, nil
}

// AddMenu implements ports.IMenuRepository.
func (m *MenuRepository) AddMenu(ctx context.Context, menuData entity.MenuORM) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.MenuRepository.AddMenu")
	defer span.Finish()

	result, err := m.MasterDB.ExecContext(ctx, queryInsertMenu, menuData.Name, menuData.Thumbnail, menuData.Price, menuData.IsActive, menuData.UserId)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

// FetchMenuById implements ports.IMenuRepository.
func (m *MenuRepository) FetchMenuById(ctx context.Context, menuId int64) ([]*entity.MenuORM, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.MenuRepository.FetchMenuById")
	defer span.Finish()

	menu_data := []*entity.MenuORM{}

	err := m.MasterDB.SelectContext(ctx, &menu_data, queryGetMenuByUserId, menuId)
	if err != nil {
		return nil, err
	}

	return menu_data, nil
}
