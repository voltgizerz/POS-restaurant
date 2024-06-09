package entity

import "time"

type CreateTokenResponse struct {
	Token     string    `json:"token"`
	TokenType string    `json:"token_type"`
	ExpiredAt time.Time `json:"expired_at"`
}

type LoginResponse struct {
	UserID    int64     `json:"user_id"`
	RoleID    int64     `json:"role_id"`
	Token     string    `json:"token"`
	TokenType string    `json:"token_type"`
	ExpiredAt time.Time `json:"expired_at"`
}
