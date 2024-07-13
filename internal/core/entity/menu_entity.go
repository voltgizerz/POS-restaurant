package entity

import (
	"time"
)

type Menu struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UserID    int64     `json:"user_id"`
	Thumbnail string    `json:"thumbnail"`
	Price     float64   `json:"price"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type MenuResponse struct {
	ID        int64   `db:"id"`
	Name      string  `db:"name"`
	Thumbnail string  `db:"thumbnail"`
	Price     float64 `db:"price"`
	IsActive  bool    `db:"is_active"`
}
