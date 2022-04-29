package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserHistoryItem struct {
	ID            primitive.ObjectID    `bson:"_id"`
	UserName      string                `bson:"username"`
	CrimeAnalysis CrimeAnalysisResponse `bson:"analysis"`
	RequestDate   time.Time             `bson:"time"`
}

//TODO: jwt и прочие признаки аутентификации
type UserHistoryRequest struct {
	ID primitive.ObjectID `bson:"_id"`
}

type UserHistoryResponse struct {
	History []UserHistoryItem `json:"history"`
}
