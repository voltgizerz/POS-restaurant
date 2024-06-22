package handler

type (
	loginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	registerRequest struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		Name            string `json:"name"`
		ConfirmPassword string `json:"confirm_password"`
		Email           string `json:"email"`
	}
)
