package handlers

import (
	"context"
	"errors"
	"githhub/iakozlov/crime-app-gateway/internal/domain"
)

var (
	ErrTimout = errors.New("the request execution timeout has expired")
)

type CrimeAnalysisService interface {
	CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest) (domain.CrimeAnalysisResponse, error)
}
