package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Menu struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	UserId    int64           `json:"user_id"`
	Thumbnail string          `json:"thumbnail"`
	Price     decimal.Decimal `json:"price"`
	IsActive  bool            `json:"is_active"`
	CreatedAt time.Time       `json:"created_at"`
}

type MenuOrm struct {
	ID        int64           `db:"id"`
	Name      string          `db:"name"`
	UserId    int64           `db:"user_id"`
	Thumbnail string          `db:"thumbnail"`
	Price     decimal.Decimal `db:"price"`
	IsActive  bool            `db:"is_active"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt time.Time       `db:"updated_at"`
}

type MenuResponse struct {
	ID        int64           `db:"id"`
	Name      string          `db:"name"`
	Thumbnail string          `db:"thumbnail"`
	Price     decimal.Decimal `db:"price"`
	IsActive  bool            `db:"is_active"`
}
