package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/voltgizerz/POS-restaurant/internal/core/models"
	"github.com/voltgizerz/POS-restaurant/internal/core/ports"
)

const (
	queryInsertMenu                    = `INSERT INTO food_menus (name,thumbnail,price,is_active,user_id) values (?,?,?,?,?)`
	queryGetMenuByUserID               = `SELECT id,name,price,thumbnail,is_active from food_menus where user_id = ? and is_active = 1`
	queryUpdateActiveBatchMenuByUserID = `UPDATE food_menus set is_active = 0 , deleted_at = ? where user_id = ?`
	queryUpdateMenuActiveByMenuID      = `UPDATE food_menus set is_active = 0 , deleted_at = ? where id = ?`
	queryUpdateMenuByMenuID            = `UPDATE food_menus set name = ? ,price = ? , thumbnail = ? ,is_active = ? , user_id = ? , updated_at = ? where id = ?`
)

type MenuRepository struct {
	MasterDB *sqlx.DB
}

func NewMenuRepository(opts RepositoryOpts) ports.IMenuRepository {
	return &MenuRepository{
		MasterDB: opts.Database.MasterDB,
	}
}

// execQuery executes a SQL query using the provided transaction or master DB.
func (m *MenuRepository) execQuery(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	if tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}

	return m.MasterDB.ExecContext(ctx, query, args...)
}

// UpdateActiveMenuBatchUser implements ports.IMenuRepository.
func (m *MenuRepository) UpdateActiveMenuBatchUser(ctx context.Context, tx *sql.Tx, userID int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.MenuRepository.UpdateActiveMenuBatchUser")
	defer span.Finish()

	result, err := m.execQuery(ctx, tx, queryUpdateActiveBatchMenuByUserID, time.Now(), userID)
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

// UpdateActiveMenuByMenuID implements ports.IMenuRepository.
func (m *MenuRepository) UpdateActiveMenuByMenuID(ctx context.Context, tx *sql.Tx, menuID int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.MenuRepository.UpdateActiveMenuByMenuID")
	defer span.Finish()

	result, err := m.execQuery(ctx, tx, queryUpdateMenuActiveByMenuID, time.Now(), menuID)
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
func (m *MenuRepository) UpdateMenuByMenuID(ctx context.Context, tx *sql.Tx, menuData *models.Menu) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.MenuRepository.UpdateMenuByMenuID")
	defer span.Finish()

	result, err := m.execQuery(ctx, tx, queryUpdateMenuByMenuID, menuData.Name, menuData.Price, menuData.Thumbnail, menuData.IsActive, menuData.UserID, time.Now(), menuData.ID)
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
func (m *MenuRepository) AddMenu(ctx context.Context, menuData models.Menu) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.MenuRepository.AddMenu")
	defer span.Finish()

	result, err := m.MasterDB.ExecContext(ctx, queryInsertMenu, menuData.Name, menuData.Thumbnail, menuData.Price, menuData.IsActive, menuData.UserID)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

// FetchMenuById implements ports.IMenuRepository.
func (m *MenuRepository) FetchMenuById(ctx context.Context, menuId int64) ([]*models.Menu, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.MenuRepository.FetchMenuById")
	defer span.Finish()

	menuData := []*models.Menu{}

	err := m.MasterDB.SelectContext(ctx, &menuData, queryGetMenuByUserID, menuId)
	if err != nil {
		return nil, err
	}

	return menuData, nil
}
