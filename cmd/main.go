package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"github.com/voltgizerz/POS-restaurant/config"
	"github.com/voltgizerz/POS-restaurant/database"
	"github.com/voltgizerz/POS-restaurant/internal/app/repository"
	"github.com/voltgizerz/POS-restaurant/pkg/jeager"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

var wg sync.WaitGroup

func main() {
	logger.Init()

	cfg := config.NewConfig()
	defer handlePanic()

	closer, err := jeager.NewJeager(cfg.App.Name)
	if err != nil {
		logger.LogStdErr.Errorf("[NewJeager] Error initializing Jaeger: %v\n", err)
	}
	defer closer.Close()

	ctx := context.Background()

	// Init database
	db := database.InitDatabase(ctx, cfg.Database)

	repositoryOpts := repository.RepositoryOpts{
		Database: db,
	}

	_ = repository.NewUserRepository(repositoryOpts)

	logger.LogStdOut.Info("Application is now running. Press CTRL-C to exit.")

	// Wait for a termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	logger.LogStdOut.Warnln("Application is exiting. Graceful shutdown in action...")
}

func handlePanic() {
	if r := recover(); r != nil {
		// Log panic location
		stack := make([]byte, 4096)
		runtime.Stack(stack, false)
		logger.LogStdErr.WithField("panic", r).WithField("stack_trace", string(stack)).Error("Panic occurred!")
	}
}
