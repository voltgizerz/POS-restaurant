package handler

type (
	loginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	registerRequest struct {
		Name            string `json:"name" validate:"required"`
		Username        string `json:"username" validate:"required"`
		Email           string `json:"email" validate:"required"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirm_password" validate:"required"`
	}
)
