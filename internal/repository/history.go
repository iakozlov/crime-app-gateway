package repository

import (
	"context"
	"github.com/iakozlov/crime-app-gateway/internal/domain"
	"github.com/iakozlov/crime-app-gateway/internal/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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
	if err != nil {
		return nil, err
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

	cursor.Close(ctx)

	return response, nil
}

func (u UserRepository) AddUserHistoryItem(ctx context.Context, item domain.UserHistoryItem, username string) error {
	historyDatabase := u.client.Database(databaseName)
	historyCollection := historyDatabase.Collection(collectionName)
	change := bson.M{"$push": bson.M{historyName: item}}

	_, err := historyCollection.UpdateOne(
		ctx,
		bson.M{collectionKey: username},
		change)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (u UserRepository) CreateUser(ctx context.Context, userName string) error {
	historyDatabase := u.client.Database(databaseName)
	historyCollection := historyDatabase.Collection(collectionName)

	_, err := historyCollection.InsertOne(ctx, bson.D{
		{Key: collectionKey, Value: userName},
		{Key: historyName, Value: bson.A{}},
	})

	if err != nil {
		log.Fatal()
	}

	return nil
}
