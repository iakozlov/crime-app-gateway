package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github/iakozlov/crime-app-gateway/cmd/config"
	"github/iakozlov/crime-app-gateway/internal/repository"
	"github/iakozlov/crime-app-gateway/internal/service"
	"github/iakozlov/crime-app-gateway/pkg/db"
)

const (
	configPath = "./config/config.yaml"
)

func main() {
	//TODO: rename to context and use for db connection
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

	fmt.Println("some message")
}
