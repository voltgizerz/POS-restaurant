package interactor

import (
	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/internal/core/ports"
)

type APInteractor struct {
	CfgAPI      config.API
	AuthHandler ports.IAuthHandler
	MenuHandler ports.IMenuHandler
}

type AuthHandler struct {
	AuthService ports.IAuthService
}

type AuthService struct {
	JWTService     ports.IJWTAuth
	TxRepository   ports.ITxRepository
	UserRepository ports.IUserRepository
}

type MenuHandler struct {
	MenuService ports.IMenuService
}

type MenuService struct {
	TxRepository   ports.ITxRepository
	MenuRepository ports.IMenuRepository
}
