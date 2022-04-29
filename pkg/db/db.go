package db

import (
	"context"
	"github/iakozlov/crime-app-gateway/cmd/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, config config.DbConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.Uri)
	client, err := mongo.Connect(ctx, clientOptions)

	return client, err
}
