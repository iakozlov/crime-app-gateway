package handlers

import (
	"context"
	"errors"
	"githhub/iakozlov/crime-app-gateway/internal/domain/analysis"
	"github.com/sirupsen/logrus"
)

var (
	ErrTimout = errors.New("the request execution timeout has expired")
)

type CrimeAnalysisService interface {
	CrimeAnalysis(ctx context.Context, request analysis.CrimeAnalysisRequest) (analysis.CrimeAnalysisResponse, error)
}

type CrimeAnalysisHandler struct {
	crimeAnalysisService CrimeAnalysisService
	log                  *logrus.Logger
}

func NewCrimeAnalysisHandler(service CrimeAnalysisService, log *logrus.Logger) *CrimeAnalysisHandler {
	return &CrimeAnalysisHandler{
		crimeAnalysisService: service,
		log:                  log,
	}
}
