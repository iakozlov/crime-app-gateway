package repository

import (
	"context"
	"github/iakozlov/crime-app-gateway/internal/domain"
	"github/iakozlov/crime-app-gateway/internal/service"
	"go.mongodb.org/mongo-driver/bson"
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
		var item domain.UserHistoryItem
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}

		results = append(results, item)
	}

	response := &domain.UserHistoryResponse{
		History: results,
	}

	cursor.Close(ctx)

	return response, nil
}

func (u UserRepository) AddUserHistoryItem(ctx context.Context, item domain.UserHistoryItem) error {
	historyDatabase := u.client.Database(databaseName)
	historyCollection := historyDatabase.Collection(collectionName)
	change := bson.M{"$push": bson.M{historyName: item}}

	_, err := historyCollection.UpdateOne(
		ctx,
		bson.M{collectionKey: item.UserName},
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
