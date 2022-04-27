package history

import (
	"githhub/iakozlov/crime-app-gateway/internal/domain/analysis"
	"time"
)

type UserHistoryItem struct {
	UserName      string                         `json:"username"`
	CrimeAnalysis analysis.CrimeAnalysisResponse `json:"analysis"`
	RequestDate   time.Time                      `json:"time"`
}
