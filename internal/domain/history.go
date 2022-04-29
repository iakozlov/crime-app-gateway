package domain

import "time"

type UserHistoryItem struct {
	UserName      string                `json:"username"`
	CrimeAnalysis CrimeAnalysisResponse `json:"analysis"`
	RequestDate   time.Time             `json:"time"`
}

//TODO: jwt и прочие признаки аутентификации
type UserHistoryRequest struct {
	UserName string `json:"username"`
}

type UserHistoryResponse struct {
	History []UserHistoryItem `json:"history"`
}
