package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github/iakozlov/crime-app-gateway/cmd/config"
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
	fmt.Println("some message")
}
