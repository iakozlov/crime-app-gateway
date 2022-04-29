package domain

import (
	"time"
)

type UserHistoryItem struct {
	UserName      string                `bson:"username"`
	CrimeAnalysis CrimeAnalysisResponse `bson:"analysis"`
	RequestDate   time.Time             `bson:"time"`
}

//TODO: jwt и прочие признаки аутентификации
type UserHistoryRequest struct {
	UserName string `bson:"username"`
}

//TODO: мб сделать словарь
type UserHistoryResponse struct {
	History []UserHistoryItem `json:"history"`
}
