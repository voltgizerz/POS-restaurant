package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/internal/adapters/api"
	"github.com/voltgizerz/POS-restaurant/internal/adapters/api/handler"
	"github.com/voltgizerz/POS-restaurant/internal/adapters/api/middleware"
	"github.com/voltgizerz/POS-restaurant/internal/adapters/auth"
	"github.com/voltgizerz/POS-restaurant/internal/adapters/repository"
	"github.com/voltgizerz/POS-restaurant/internal/core/interactor"
	"github.com/voltgizerz/POS-restaurant/internal/core/service"
	"github.com/voltgizerz/POS-restaurant/internal/database"
	"github.com/voltgizerz/POS-restaurant/pkg/jeager"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

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

	// Initialize Middleware
	jwtMiddleware := middleware.NewJWTAuthMiddleware(authJWT)

	repoOpts := repository.RepositoryOpts{
		Database: db,
	}

	// Initialize Repositories
	txRepo := repository.NewTxRepository(repoOpts)
	userRepo := repository.NewUserRepository(repoOpts)
	menuRepo := repository.NewMenuRepository(repoOpts)

	// Initialize Services
	authService := service.NewAuthService(interactor.AuthService{
		JWTService:     authJWT,
		TxRepository:   txRepo,
		UserRepository: userRepo,
	})

	menuService := service.NewMenuService(interactor.MenuService{
		TxRepository:   txRepo,
		MenuRepository: menuRepo,
	})

	// Initialize Handlers
	authHandler := handler.NewAuthHandler(interactor.AuthHandler{
		AuthService: authService,
	})

	menuHandler := handler.NewMenuHandler(interactor.MenuHandler{
		MenuService: menuService,
	})

	interactoAPI := interactor.APInteractor{
		CfgAPI:      cfg.API,
		AuthHandler: authHandler,
		MenuHandler: menuHandler,
	}

	// Start API server
	go startAPIServer(interactoAPI, jwtMiddleware)
	logger.Log.Info("Application is now running. Press CTRL-C to exit.")

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

func startAPIServer(interactor interactor.APInteractor, jwtMiddleware middleware.JWTAuth) {
	httpServer := api.NewServer(interactor, jwtMiddleware)
	httpServer.Initialize()
}

func waitForSignal() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	logger.Log.Warnln("Application is exiting. Graceful shutdown in action...")
}
