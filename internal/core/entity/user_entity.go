package entity

import "time"

type User struct {
	ID        int64     `json:"id"`
	RoleID    int64     `json:"role_id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
type RegisterResponse struct {
	UserID int64 `json:"user_id"`
}
