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
		ConfirmPassword string `json:"confirm_password"`
		Email           string `json:"email"`
	}

	addMenuRequest struct {
		Name      string `json:"name"`
		Thumbnail string `json:"thumbnail"`
		Price     string `json:"price"`
		UserId    string `json:"userid"`
	}

	updateMenuRequest struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Thumbnail string `json:"thumbnail"`
		Price     string `json:"price"`
		UserId    string `json:"userid"`
		IsActive  string `json:"isactive"`
	}

	getMenuRequest struct {
		UserId string `json:"userid"`
	}

	menuIdRequest struct {
		MenuId string `json:"menuid"`
	}
)
