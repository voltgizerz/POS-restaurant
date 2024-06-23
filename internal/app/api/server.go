package api

import (
	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/internal/app/api/middleware"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

type Server struct {
	cfgAPI        config.API
	jwtMiddleware middleware.JWTAuth
	authHandler   ports.IAuthHandler
	menuHandler   ports.IMenuHandler
}

func NewServer(interactor interactor.APInteractor, jwtMiddleware middleware.JWTAuth) *Server {
	return &Server{
		cfgAPI:        interactor.CfgAPI,
		jwtMiddleware: jwtMiddleware,
		authHandler:   interactor.AuthHandler,
		menuHandler:   interactor.MenuHandler,
	}
}

func (s *Server) Initialize() {
	app := s.InitRouter()

	err := app.Listen(":" + s.cfgAPI.PORT)
	if err != nil {
		logger.LogStdErr.Fatalf("[Initialize] error on: %s", err.Error())
	}
}
