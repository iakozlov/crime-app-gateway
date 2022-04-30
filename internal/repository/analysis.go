package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github/iakozlov/crime-app-gateway/internal/domain"
	"github/iakozlov/crime-app-gateway/internal/service"
	"io/ioutil"
	"net/http"
)

const (
	methodType = "POST"
	uri        = "localhost:5000"
)

type CrimeAnalysisRepository struct {
	client http.Client
}

func NewCrimeAnalysisRepository(client http.Client) service.CrimeAnalysisRepository {
	return CrimeAnalysisRepository{
		client: client,
	}
}

func (r CrimeAnalysisRepository) CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest) (domain.CrimeAnalysisResponse, error) {
	//TODO: разобраться точно с полями в теле запроса
	requestBody, err := json.Marshal(map[string]string{
		"X":    request.Lat,
		"Y":    request.Lng,
		"Date": request.Date,
	})
	if err != nil {
		fmt.Errorf("can't make request body, err: %w", err)
	}

	httpRequest, err := http.NewRequest(methodType, uri, bytes.NewBuffer(requestBody))
	httpRequest.Header.Set("Content-type", "application/json")
	if err != nil {
		fmt.Errorf("can't make request, err: %w", err)
	}

	response, err := r.client.Do(httpRequest)
	if err != nil {
		fmt.Errorf("can't get data from crime service, err: %w", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("can't read response body, err: %w", err)
	}

	var crimes []domain.CrimeInfoModel

	err = json.Unmarshal(body, &crimes)
	if err != nil {
		fmt.Errorf("can't decode response body, err: %w", err)
	}

	result := domain.CrimeAnalysisResponse{
		Crimes: crimes,
	}

	return result, nil
}
