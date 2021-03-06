package db

import (
	"context"
	"github.com/iakozlov/crime-app-gateway/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, config config.DbConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.Uri)
	client, err := mongo.Connect(ctx, clientOptions)

	return client, err
}
