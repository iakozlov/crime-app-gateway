package main

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github/iakozlov/crime-app-gateway/config"
	"github/iakozlov/crime-app-gateway/internal/handlers"
	"github/iakozlov/crime-app-gateway/internal/repository"
	"github/iakozlov/crime-app-gateway/internal/service"
	"github/iakozlov/crime-app-gateway/pkg/db"
	"github/iakozlov/crime-app-gateway/pkg/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	configPath = "/Users/ivkozlov/Desktop/Coding/crime-app-gateway/config/config.yaml"
)

// @title Crime app auth
// @version 1.0
// @description Crime app auth provides authentication for crime-app microservices.

// @host localhost:8000
// @BasePath /
func main() {
	ctx := context.Background()
	log := logrus.New()

	cfg, error := config.Read(configPath)
	if error != nil {
		log.Fatal("can't load config, err: %w", error)
	}

	mongoClient, err := db.Connect(ctx, cfg.DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	historyRepository := repository.NewUserRepository(mongoClient)
	historyService := service.NewUserHistoryService(historyRepository)

	client := http.Client{
		Timeout: cfg.CtxTimeout,
	}

	analysisRepository := repository.NewCrimeAnalysisRepository(client)
	analysisService := service.NewCrimeAnalysisService(analysisRepository)

	e := echo.New()
	handlers.InitCommonRoutes(e)
	handler := handlers.NewCrimeAnalysisHandler(analysisService, historyService, log)
	handler.InitRoutes(e, cfg.CtxTimeout)

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
