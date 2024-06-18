package entity

import "time"

type User struct {
	ID        int64     `json:"id"`
	RoleID    int64     `json:"role_id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password_hashed"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type UserORM struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password_hashed"`
	IsActive  bool      `db:"is_active"`
	RoleID    int64     `db:"role_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type RegisterResponse struct {
	User_id int64 `json:"user_id"`
}
