package entity

import "time"

type (
	LoginRequest struct {
		Username string `json:"username" validate:"required,alphanum"`
		Password string `json:"password" validate:"required,alphanumunicode"`
	}

	RegisterRequest struct {
		Name            string `json:"name" validate:"required,alphanum"`
		Username        string `json:"username" validate:"required,alphanum"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,alphanumunicode"`
		ConfirmPassword string `json:"confirm_password" validate:"required,alphanumunicode"`
	}
)

type (
	LoginResponse struct {
		UserID    int64     `json:"user_id"`
		RoleID    int64     `json:"role_id"`
		Token     string    `json:"token"`
		TokenType string    `json:"token_type"`
		ExpiredAt time.Time `json:"expired_at"`
	}

	CreateTokenResponse struct {
		Token     string    `json:"token"`
		TokenType string    `json:"token_type"`
		ExpiredAt time.Time `json:"expired_at"`
	}
)
