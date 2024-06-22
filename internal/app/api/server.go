package api

import (
	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

type Server struct {
	cfgAPI      config.API
	userHandler ports.IUserHandler
	menuHandler ports.IMenuHandler
}

func NewServer(interactor interactor.APInteractor) *Server {
	return &Server{
		cfgAPI:      interactor.CfgAPI,
		userHandler: interactor.UserHandler,
		menuHandler: interactor.MenuHandler,
	}
}

func (s *Server) Initialize() {
	if s.menuHandler == nil {
		print("Have Handler")
	}
	app := s.InitRouter()

	err := app.Listen(":" + s.cfgAPI.PORT)
	if err != nil {
		logger.LogStdErr.Fatalf("[Initialize] error on: %s", err.Error())
	}
}
