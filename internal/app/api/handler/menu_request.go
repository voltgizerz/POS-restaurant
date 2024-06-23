package handler

type (
	addMenuRequest struct {
		Name      string `json:"name"`
		Thumbnail string `json:"thumbnail"`
		Price     string `json:"price"`
		UserID    int64  `json:"user_id"`
		IsActive  string `json:"is_active"`
	}

	updateMenuRequest struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		Thumbnail string `json:"thumbnail"`
		Price     string `json:"price"`
		UserID    int64  `json:"user_id"`
		IsActive  string `json:"is_active"`
	}
)
