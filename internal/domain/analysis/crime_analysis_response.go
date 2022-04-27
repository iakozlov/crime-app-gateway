package analysis

type CrimeAnalysisResponse struct {
	Crimes []CrimeInfoModel `json:"crimes"`
}
