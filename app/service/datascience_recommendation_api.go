package service

import (
	"bytes"
	"encoding/json"
	"github.com/ekkinox/fx-template/modules/fxconfig"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"io"
	"net/http"
)

// DataScienceRecommendationApi Handles the communication with the data science recommendation API.
type DataScienceRecommendationApi struct {
	apiUrl string
	apiKey string

	logger *fxlogger.Logger
}

// NewDataScienceRecommendationApi Creates a new DataScienceRecommendationApi.
func NewDataScienceRecommendationApi(config *fxconfig.Config, logger *fxlogger.Logger) *DataScienceRecommendationApi {
	return &DataScienceRecommendationApi{
		apiUrl: config.GetString("config.datascience-api.url"),
		apiKey: config.GetString("config.datascience-api.key"),
		logger: logger,
	}
}

// GetRecommendationsByEntityAndType Gets recommendations by entity and type.
func (s *DataScienceRecommendationApi) GetRecommendationsByEntityAndType() (string, error) {
	s.logger.Debug().Msg("Call to get recommendations by entity and type")
	s.logger.Debug().Msgf("url: %v", s.apiUrl)
	s.logger.Debug().Msgf("key: %v", s.apiKey)

	url := s.apiUrl + "/retailers_recommendations/v5/reco_by_type"
	// TODO: make it dynamic
	body, _ := json.Marshal(map[string]any{
		"ids":  []int{18},
		"type": 2,
	})

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		s.logger.Error().Msgf("DS API HTTP error: %v", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", s.apiKey)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Error().Msgf("DS API HTTP exec error: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error().Msgf("DS API read body error: %v", err)
		return "", err
	}

	bodyStr := string(respBody)
	s.logger.Debug().Msgf("response: %v", bodyStr)

	return bodyStr, nil
}
