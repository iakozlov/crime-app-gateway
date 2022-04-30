package service

import (
	"context"
	"github/iakozlov/crime-app-gateway/internal/domain"
	"github/iakozlov/crime-app-gateway/internal/handlers"
)

type CrimeAnalysisRepository interface {
	CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest) (domain.CrimeAnalysisResponse, error)
}

type CrimeAnalysisService struct {
	crimeAnalysisRepo CrimeAnalysisRepository
}

func NewCrimeAnalysisService(repo CrimeAnalysisRepository) handlers.CrimeAnalysisService {
	return CrimeAnalysisService{
		crimeAnalysisRepo: repo,
	}
}

func (s CrimeAnalysisService) CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest) (domain.CrimeAnalysisResponse, error) {
	return s.crimeAnalysisRepo.CrimeAnalysis(ctx, request)
}
