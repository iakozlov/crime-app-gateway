package service

import (
	"context"
	"github.com/iakozlov/crime-app-gateway/internal/domain"
	"github.com/iakozlov/crime-app-gateway/internal/handlers"
)

type CrimeAnalysisRepository interface {
	CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest) (*domain.CrimeAnalysisResponse, error)
}

type CrimeAnalysisService struct {
	crimeAnalysisRepo CrimeAnalysisRepository
	historyService    handlers.UserHistoryService
}

func NewCrimeAnalysisService(repo CrimeAnalysisRepository, historyService handlers.UserHistoryService) handlers.CrimeAnalysisService {
	return CrimeAnalysisService{
		crimeAnalysisRepo: repo,
		historyService:    historyService,
	}
}

func (s CrimeAnalysisService) CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest) (*domain.CrimeAnalysisResponse, error) {
	crimeAnalysis, err := s.crimeAnalysisRepo.CrimeAnalysis(ctx, request)
	if err != nil {
		return nil, err
	}

	history := domain.UserHistoryItem{
		CrimeAnalysis: *crimeAnalysis,
		RequestDate:   request.Date,
		Address:       request.Address,
	}
	err = s.historyService.AddHistory(ctx, history, request.UserName)

	if err != nil {
		return nil, err
	}

	return crimeAnalysis, nil
}
