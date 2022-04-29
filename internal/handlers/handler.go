package handlers

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github/iakozlov/crime-app-gateway/internal/domain"
)

var (
	ErrTimout = errors.New("the request execution timeout has expired")
)

type CrimeAnalysisService interface {
	CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest) (domain.CrimeAnalysisResponse, error)
}

type UserHistoryService interface {
	History(ctx context.Context, request domain.UserHistoryRequest) (*domain.UserHistoryResponse, error)
}

type CrimeAnalysisHandler struct {
	crimeAnalysisService CrimeAnalysisService
	userHistoryService   UserHistoryService
	log                  *logrus.Logger
}

func NewCrimeAnalysisHandler(analysisService CrimeAnalysisService, historyService UserHistoryService, log *logrus.Logger) *CrimeAnalysisHandler {
	return &CrimeAnalysisHandler{
		crimeAnalysisService: analysisService,
		userHistoryService:   historyService,
		log:                  log,
	}
}
