package service

import (
	"context"
	"github/iakozlov/crime-app-gateway/internal/domain"
)

type CrimeAnalysisService struct {
}

func (s CrimeAnalysisService) CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest) (domain.CrimeAnalysisResponse, error) {
	return domain.CrimeAnalysisResponse{}, nil
}
