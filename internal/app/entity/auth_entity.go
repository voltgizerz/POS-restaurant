package entity

import "time"

type CreateTokenResponse struct {
	Token     string    `json:"token"`
	TokenType string    `json:"token_type"`
	ExpiredAt time.Time `json:"expired_at"`
}