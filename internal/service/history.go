package service

import (
	"context"
	"fmt"
	"github/iakozlov/crime-app-gateway/internal/domain"
	"github/iakozlov/crime-app-gateway/internal/handlers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHistoryRepository interface {
	UserHistory(ctx context.Context, id primitive.ObjectID) (*domain.UserHistoryResponse, error)
	AddUserHistoryItem(ctx context.Context, item domain.UserHistoryItem) error
}

type UserHistoryService struct {
	historyRepo UserHistoryRepository
}

func NewUserHistoryService(rep UserHistoryRepository) handlers.UserHistoryService {
	return &UserHistoryService{
		historyRepo: rep,
	}
}

func (u UserHistoryService) History(ctx context.Context, request domain.UserHistoryRequest) (*domain.UserHistoryResponse, error) {
	storedHistory, err := u.historyRepo.UserHistory(ctx, request.ID)
	if err != nil {
		return nil, fmt.Errorf("can't get user history from database, err: %w", err)
	}

	return storedHistory, nil
}
