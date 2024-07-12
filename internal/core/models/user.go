package models

import "time"

type User struct {
	ID             int64     `db:"id"`
	Name           string    `db:"name"`
	Username       string    `db:"username"`
	Email          string    `db:"email"`
	PasswordHashed string    `db:"password_hashed"`
	IsActive       bool      `db:"is_active"`
	RoleID         int64     `db:"role_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
