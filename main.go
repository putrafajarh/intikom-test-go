package main

import (
	"context"
	"os/signal"
	"syscall"

	"intikom-test-go/config"
	"intikom-test-go/database"
	"intikom-test-go/router"
	"intikom-test-go/utils"

	"github.com/gin-contrib/graceful"
	"go.uber.org/zap"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := utils.InitLogger()
	defer logger.Sync()

	if err := config.InitConfig(); err != nil {
		logger.Fatal("Failed to initialize config", zap.Error(err))
	}
	masterDSN, replicaDSN := config.DatabaseConfig()
	if err := database.DBConnection(masterDSN, replicaDSN); err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	gracefulGin, err := graceful.New(router.InitRouter())
	if err != nil {
		logger.Fatal("Failed to initialize graceful", zap.Error(err))
	}
	defer gracefulGin.Close()

	if err := gracefulGin.RunWithContext(ctx); err != nil && err != context.Canceled {
		panic(err)
	}
}
