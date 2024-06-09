package handler

type (
	loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	registerRequest struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		Name            string `json:"name"`
		ConfirmPassword string `json:"confirmpassword"`
		Email           string `json:"email"`
	}
)
