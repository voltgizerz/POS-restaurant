package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/database"
	"github.com/voltgizerz/POS-restaurant/internal/app/api"
	"github.com/voltgizerz/POS-restaurant/internal/app/api/auth"
	"github.com/voltgizerz/POS-restaurant/internal/app/api/handler"
	"github.com/voltgizerz/POS-restaurant/internal/app/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/app/repository"
	"github.com/voltgizerz/POS-restaurant/internal/app/service"
	"github.com/voltgizerz/POS-restaurant/pkg/jeager"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

var wg sync.WaitGroup

func main() {
	logger.Init()

	cfg := config.NewConfig()
	defer handlePanic()

	// Initialize Jaeger
	initJaeger(cfg.App.Name)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize database
	db := database.InitDatabase(ctx, cfg.Database)

	// Initialize Auth JWT
	authJWT := auth.NewAuthJWT(cfg.API.JWTSecretKey)

	repoOpts := repository.RepositoryOpts{
		Database: db,
	}

	// Initialize Repositories
	userRepo := repository.NewUserRepository(repoOpts)

	menuRepo := repository.NewMenuRepository(repoOpts)

	// Initialize Services
	userService := service.NewUserService(interactor.UserService{
		AuthService:    authJWT,
		UserRepository: userRepo,
	})

	menuService := service.NewMenuService(interactor.MenuService{
		AuthService:    authJWT,
		MenuRepository: menuRepo,
	})

	// Initialize Handlers
	userHandler := handler.NewUserHandler(interactor.UserHandler{
		UserService: userService,
	})

	menuHandler := handler.NewMenuHandler(interactor.MenuHandler{
		MenuService: menuService,
		AuthService: authJWT,
	})

	interactoAPI := interactor.APInteractor{
		CfgAPI:      cfg.API,
		UserHandler: userHandler,
		MenuHandler: menuHandler,
	}

	// Start API server
	go startAPIServer(interactoAPI)

	// Wait for termination signal
	waitForSignal()
}

func initJaeger(serviceName string) {
	closer, err := jeager.NewJeager(serviceName)
	if err != nil {
		logger.LogStdErr.Errorf("[NewJeager] Error initializing Jaeger: %v", err)
		return
	}
	defer closer.Close()
}

func handlePanic() {
	if r := recover(); r != nil {
		logger.LogStdErr.Errorf("Panic occurred: %v", r)
	}
}

func startAPIServer(interactor interactor.APInteractor) {
	httpServer := api.NewServer(interactor)
	httpServer.Initialize()
}

func waitForSignal() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	logger.Log.Warnln("Application is exiting. Graceful shutdown in action...")
}
