package domain

type CrimeAnalysisRequest struct {
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Address string `json:"address"`
}

type CrimeInfoModel struct {
	Name        string  `json:"name"`
	Probability float64 `json:"probability"`
}

type CrimeAnalysisResponse struct {
	Crimes map[string]map[string]string `json:"crimes"`
}
