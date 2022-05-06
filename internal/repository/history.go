package repository

import (
	"context"
	"fmt"
	"github.com/iakozlov/crime-app-gateway/internal/domain"
	"github.com/iakozlov/crime-app-gateway/internal/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName   = "historyCollector"
	collectionName = "userHistories"
	collectionKey  = "username"
	historyName    = "history"
)

type UserRepository struct {
	client *mongo.Client
}

type UserHistoryRecord struct {
	ID      primitive.ObjectID       `bson:"_id"`
	UseName string                   `bson:"username"`
	History []domain.UserHistoryItem `bson:"history"`
}

func NewUserRepository(client *mongo.Client) service.UserHistoryRepository {
	return &UserRepository{
		client: client,
	}
}

func (u UserRepository) UserHistory(ctx context.Context, userName string) (*domain.UserHistoryResponse, error) {
	historyDatabase := u.client.Database(databaseName)
	historyCollection := historyDatabase.Collection(collectionName)

	var results []domain.UserHistoryItem
	cursor, err := historyCollection.Find(ctx, bson.M{collectionKey: userName})
	defer cursor.Close(ctx)

	if err != nil {
		return nil, fmt.Errorf("can't find history in db, err: %w", err)
	}
	for cursor.Next(ctx) {
		var record UserHistoryRecord
		err := cursor.Decode(&record)
		if err != nil {
			return nil, err
		}

		results = record.History
	}

	response := &domain.UserHistoryResponse{
		History: results,
	}

	return response, nil
}

func (u UserRepository) AddUserHistoryItem(ctx context.Context, item domain.UserHistoryItem, username string) error {
	historyDatabase := u.client.Database(databaseName)
	historyCollection := historyDatabase.Collection(collectionName)
	change := bson.M{"$push": bson.M{historyName: item}}

	if _, err := historyCollection.UpdateOne(ctx, bson.M{collectionKey: username}, change); err != nil {
		return fmt.Errorf("can't add user history to database, err: %w", err)
	}

	return nil
}

func (u UserRepository) CreateUser(ctx context.Context, userName string) error {
	historyDatabase := u.client.Database(databaseName)
	historyCollection := historyDatabase.Collection(collectionName)

	insertedInfo := bson.D{{
		Key: collectionKey, Value: userName},
		{Key: historyName, Value: bson.A{}}}

	if _, err := historyCollection.InsertOne(ctx, insertedInfo); err != nil {
		return fmt.Errorf("can't create user in database, err: %w", err)
	}

	return nil
}
