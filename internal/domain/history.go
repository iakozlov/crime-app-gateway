package domain

type UserHistoryItem struct {
	CrimeAnalysis CrimeAnalysisResponse `bson:"analysis"`
	RequestDate   string                `bson:"time"`
	Address       string                `bson:"address"`
}

//TODO: jwt и прочие признаки аутентификации
type UserHistoryRequest struct {
	UserName string `bson:"username"`
}

//TODO: мб сделать словарь
type UserHistoryResponse struct {
	History []UserHistoryItem `json:"history"`
}
