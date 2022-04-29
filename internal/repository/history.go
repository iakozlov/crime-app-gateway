package repository

import (
	"context"
	"github/iakozlov/crime-app-gateway/internal/domain"
	"github/iakozlov/crime-app-gateway/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName   = "historyCollector"
	collectionName = "userHistories"
)

type UserRepository struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) service.UserHistoryRepository {
	return &UserRepository{
		client: client,
	}
}

func (u UserRepository) UserHistory(ctx context.Context, id primitive.ObjectID) (*domain.UserHistoryResponse, error) {
	return nil, nil
}

func (u UserRepository) AddUserHistoryItem(ctx context.Context, item domain.UserHistoryItem) error {
	return nil
}

func (u UserRepository) CreateUser(ctx context.Context, id primitive.ObjectID) error {
	return nil
}
