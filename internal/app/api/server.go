package api

import (
	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/ports"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

type Server struct {
	cfg         config.API
	userHandler ports.IUserHandler
}

func NewServer(interactor interactor.APInteractor) *Server {
	return &Server{
		cfg:         interactor.Cfg,
		userHandler: interactor.UserHandler,
	}
}

func (s *Server) Initialize() {
	app := s.InitRouter()

	err := app.Listen(":" + s.cfg.PORT)
	if err != nil {
		logger.LogStdErr.Fatalf("[Initialize] error on: %s", err.Error())
	}
}
