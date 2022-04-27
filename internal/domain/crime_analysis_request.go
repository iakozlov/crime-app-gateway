package domain

type CrimeAnalysisRequest struct {
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Address string `json:"address"`
}
