package domain

type CrimeAnalysisResponse struct {
	Crimes []CrimeInfoModel `json:"crimes"`
}
