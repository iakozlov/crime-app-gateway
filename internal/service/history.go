package service

import (
	"context"
	"fmt"
	"github.com/iakozlov/crime-app-gateway/internal/domain"
	"github.com/iakozlov/crime-app-gateway/internal/handlers"
)

type UserHistoryRepository interface {
	UserHistory(ctx context.Context, userName string) (*domain.UserHistoryResponse, error)
	AddUserHistoryItem(ctx context.Context, item domain.UserHistoryItem, username string) error
	CreateUser(ctx context.Context, userName string) error
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
	storedHistory, err := u.historyRepo.UserHistory(ctx, request.UserName)
	if err != nil {
		return nil, fmt.Errorf("can't get user history from database, err: %w", err)
	}

	return storedHistory, nil
}

func (u UserHistoryService) AddHistory(ctx context.Context, request domain.UserHistoryItem, username string) error {
	storedHistory, err := u.historyRepo.UserHistory(ctx, username)
	if err != nil {
		return fmt.Errorf("can't get user history from database, err: %w", err)
	}

	if len(storedHistory.History) == 0 {
		err := u.historyRepo.CreateUser(ctx, username)
		if err != nil {
			return fmt.Errorf("can't create user history in database, err: %w", err)
		}

		err = u.historyRepo.AddUserHistoryItem(ctx, request, username)
		if err != nil {
			return fmt.Errorf("can't add user hostory to database, err: %w", err)
		}

		return nil
	}

	err = u.historyRepo.AddUserHistoryItem(ctx, request, username)
	if err != nil {
		return fmt.Errorf("can't add user hostory to database, err: %w", err)
	}

	return nil

}
