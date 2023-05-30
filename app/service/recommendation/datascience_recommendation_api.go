package recommendationService

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ekkinox/fx-template/app/enum"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"io"
	"net/http"
	"strconv"
)

// DataScienceRecommendationApi Handles the communication with the data science recommendation API.
type DataScienceRecommendationApi struct {
	apiUrl string
	apiKey string

	logger        *fxlogger.Logger
	apiUrlService ApiUrl
}

// GetRecommendationsByEntityAndType Gets recommendations by entity and type.
func (s *DataScienceRecommendationApi) GetRecommendationsByEntityAndType(
	recommendableId int,
	recommendableType string,
	recommendationTypeId int,
	metadata map[string]any,
) ([]int, error) {
	// Perform the request
	response, err := s.request(recommendableId, recommendableType, recommendationTypeId, metadata)
	if err != nil {
		s.logger.Error().Msgf("DS API request error: %v", err)
		return []int{}, err
	}

	// Parse the response
	result := make(map[string]map[string][]int)
	err2 := json.Unmarshal([]byte(response), &result)
	if err2 != nil {
		s.logger.Error().Msgf("DS API JSON unmarshaling error: %v", err)
		return []int{}, err
	}

	// Get the recommendations
	recommendations, ok := result["reco"][strconv.Itoa(recommendableId)]
	if !ok {
		s.logger.Error().Msgf("DS API response error: %v", err)
		return []int{}, errors.New("DS API response error")
	}

	return recommendations, nil
}

// request Performs the request to the DataScience API.
func (s *DataScienceRecommendationApi) request(
	recommendableId int,
	recommendableType string,
	recommendationTypeId int,
	metadata map[string]any,
) (string, error) {
	url, err := s.apiUrlService.Url(recommendableType, recommendationTypeId)
	if err != nil {
		s.logger.Error().Msg("DS API URL error")
		return "", err
	}

	// Get payload
	payload, err := s.getPayload(recommendableId, recommendationTypeId, metadata)
	if err != nil {
		s.logger.Error().Msg("DS API payload error")
		return "", err
	}

	// Prepare the request payload
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		s.logger.Error().Msgf("DS API JSON payload marshaling error")
		return "", err
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonPayload))
	if err != nil {
		s.logger.Error().Msgf("DS API HTTP error")
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", s.apiKey)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Error().Msgf("DS API HTTP exec error")
		return "", err
	}
	defer resp.Body.Close()

	// Check HTTP status
	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("DS API HTTP status error: %v", resp.StatusCode)
		return "", errors.New(msg)
	}

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error().Msgf("DS API read body error")
		return "", err
	}

	bodyStr := string(respBody)

	return bodyStr, nil
}

// getPayload Gets the payload for the request.
func (s *DataScienceRecommendationApi) getPayload(
	recommendableId int,
	recommendationTypeId int,
	metadata map[string]any,
) (map[string]any, error) {
	// Category products recommendation payload
	if recommendationTypeId == enum.RetailerCategoryProductsYouMayLike {
		// Check that metadata contains category_id
		if _, ok := metadata["category_id"]; !ok {
			return nil, errors.New("missing category_id in metadata")
		}

		return map[string]any{
			"id":       recommendableId,
			"type":     recommendationTypeId,
			"category": metadata["category_id"],
		}, nil
	}

	// Generic recommendation payload
	return map[string]any{
		"ids":  []int{recommendableId},
		"type": recommendationTypeId,
	}, nil
}
