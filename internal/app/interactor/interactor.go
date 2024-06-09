package interactor

import (
	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

type APInteractor struct {
	CfgAPI      config.API
	UserHandler ports.IUserHandler
}

type UserHandler struct {
	Auth           ports.IAuth
	UserService ports.IUserService
}
