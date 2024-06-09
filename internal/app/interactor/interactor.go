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
	UserService ports.IUserService
}

type UserService struct {
	AuthService    ports.IAuth
	UserRepository ports.IUserRepository
}
