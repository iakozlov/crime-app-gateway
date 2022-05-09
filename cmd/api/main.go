package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/iakozlov/crime-app-gateway/config"
	"github.com/iakozlov/crime-app-gateway/internal/handlers"
	"github.com/iakozlov/crime-app-gateway/internal/repository"
	"github.com/iakozlov/crime-app-gateway/internal/service"
	"github.com/iakozlov/crime-app-gateway/pkg/db"
	"github.com/iakozlov/crime-app-gateway/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	configPath = "./config/config.yaml"
)

// @title Crime app auth
// @version 1.0
// @description Crime app auth provides authentication for crime-app microservices.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name CrimeJWT

// @host localhost:8000
// @BasePath /
func main() {
	ctx := context.Background()
	log := logrus.New()

	cfg, err := config.Read(configPath)
	if err != nil {
		log.Fatal("can't load config, err: %w", err)
	}

	mongoClient, err := db.Connect(ctx, cfg.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	historyRepository := repository.NewUserRepository(mongoClient)
	historyService := service.NewUserHistoryService(historyRepository)

	analysisRepository := repository.NewCrimeAnalysisRepository(
		http.Client{
			Timeout: cfg.CtxTimeout,
		})
	analysisService := service.NewCrimeAnalysisService(analysisRepository, historyService)

	e := echo.New()
	handlers.InitCommonRoutes(e)
	handler := handlers.NewCrimeAnalysisHandler(analysisService, historyService, log)
	handler.InitRoutes(e, cfg.CtxTimeout, cfg.SecretJWT)

	srv := server.NewServer(cfg.SrvConfig, e)

	// gracefully shutdown
	go func() {
		if err = srv.Run(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("running server is failed, err: %q", err)
			}
		}
	}()

	log.Info("server is running...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("server is shutting down...")

	if err = srv.Stop(ctx); err != nil {
		log.Fatalf("graceful shutdown is broken, err: %q", err)
	}
	log.Info("cmd was gracefully shut down.")
}
