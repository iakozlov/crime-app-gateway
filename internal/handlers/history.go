package handlers

import (
	"context"
	"githhub/iakozlov/crime-app-gateway/internal/domain/history"
	"github.com/sirupsen/logrus"
)

type UserHistoryService interface {
	History(ctx context.Context, request history.UserHistoryRequest) (history.UserHistoryResponse, error)
}

type UserHistoryHandler struct {
	userHistoryService UserHistoryService
	log                *logrus.Logger
}

func NewUserHistoryHandler(service UserHistoryService, log *logrus.Logger) *UserHistoryHandler {
	return &UserHistoryHandler{
		userHistoryService: service,
		log:                log,
	}
}
