package handler

type (
	loginRequest struct {
		Username string `json:"username" validate:"required,alphanum"`
		Password string `json:"password" validate:"required,alphanumunicode"`
	}

	registerRequest struct {
		Name            string `json:"name" validate:"required,alphanum"`
		Username        string `json:"username" validate:"required,alphanum"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,alphanumunicode"`
		ConfirmPassword string `json:"confirm_password" validate:"required,alphanumunicode"`
	}
)
