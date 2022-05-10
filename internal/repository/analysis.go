package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/iakozlov/crime-app-gateway/internal/domain"
	"github.com/iakozlov/crime-app-gateway/internal/service"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	methodType = "POST"
	uri        = "http://192.168.0.99:80/predict"
)

type CrimeAnalysisRepository struct {
	client http.Client
}

func NewCrimeAnalysisRepository(client http.Client) service.CrimeAnalysisRepository {
	return CrimeAnalysisRepository{
		client: client,
	}
}

func (r CrimeAnalysisRepository) CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest) (*domain.CrimeAnalysisResponse, error) {
	//TODO: разобраться точно с полями в теле запроса
	requestBody, err := json.Marshal(map[string]string{
		"X":    request.Lat,
		"Y":    request.Lng,
		"date": request.Date,
	})
	if err != nil {
		fmt.Errorf("can't make request body, err: %w", err)
	}

	httpRequest, err := http.NewRequest(methodType, uri, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Errorf("can't make request, err: %w", err)
	}

	httpRequest.Header.Set("Content-type", "application/json")

	response, err := r.client.Do(httpRequest)
	if err != nil {
		fmt.Errorf("can't get data from crime service, err: %w", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("can't read response body, err: %w", err)
	}

	var crimesDict map[string]map[string]string

	err = json.Unmarshal(body, &crimesDict)
	if err != nil {
		fmt.Errorf("can't decode response body, err: %w", err)
	}

	crimes := crimesDict["prediction"]

	var crimesInfoSlice []domain.CrimeInfoModel

	for key, value := range crimes {
		probability, err := strconv.ParseFloat(value, 3)
		if err != nil {
			return nil, fmt.Errorf("can't convert probability to float, err: %w", err)
		}

		crimesInfoSlice = append(crimesInfoSlice, domain.CrimeInfoModel{
			Name:        key,
			Probability: probability,
		})
	}

	result := &domain.CrimeAnalysisResponse{
		Crimes: crimesInfoSlice,
	}
	return result, nil
}
