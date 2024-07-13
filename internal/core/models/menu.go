package models

import "time"

type MenuORM struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	UserID    int64     `db:"user_id"`
	Thumbnail string    `db:"thumbnail"`
	Price     float64   `db:"price"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
