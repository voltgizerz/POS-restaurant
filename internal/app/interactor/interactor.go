package interactor

import (
	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
)

type APInteractor struct {
	CfgAPI      config.API
	UserHandler ports.IUserHandler
	MenuHandler ports.IMenuHandler
}

type UserHandler struct {
	UserService ports.IUserService
}

type MenuHandler struct {
	MenuService ports.IMenuService
	AuthService ports.IAuth
}

type UserService struct {
	AuthService    ports.IAuth
	UserRepository ports.IUserRepository
}

type MenuService struct {
	AuthService    ports.IAuth
	MenuRepository ports.IMenuRepository
}
