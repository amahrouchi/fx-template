package recommendationService

import (
	"encoding/json"
	"fmt"
	"github.com/ekkinox/fx-template/app/service/cache"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// RecommendationClient Gather recommendations from the recommendation API.
type RecommendationClient struct {
	ttl int

	recommendationApi RecommendationApiContract
	cacheService      cacheService.CacheContract
	logger            *fxlogger.Logger
}

// GetRecommendationsByEntityAndType Get recommendations by entity and type.
func (rc *RecommendationClient) GetRecommendationsByEntityAndType(
	recommendableId int,
	recommendableType string,
	recommendationTypeId int,
) ([]int, error) {
	key := fmt.Sprintf("{%s}.{%d}.{%d}", recommendableType, recommendableId, recommendationTypeId)

	// Get cached recommendations
	cachedRecos, err := rc.cacheService.Get(key)
	if err != nil {
		return []int{}, err
	}

	// No cached recommendations found
	if cachedRecos == "" {
		rc.logger.Info().Msgf("No cached recommendations found for key: %s", key)

		// Get recommendations from API
		apiRecos, err := rc.recommendationApi.GetRecommendationsByEntityAndType(recommendableId, recommendableType, recommendationTypeId, map[string]any{})
		if err != nil {
			rc.logger.Err(err).Msgf("Failed to get recommendations from the API for key: %s", key)
			return []int{}, err
		}

		rc.logger.Info().Msgf("Got recommendations from the API for key=%s, recos=%v", key, apiRecos)

		// JSONify recommendations
		jsonRecos, err := json.Marshal(apiRecos)
		if err != nil {
			rc.logger.Err(err).Msgf("Failed to jsonify recommendations from the API for key: %s", key)
			return []int{}, err
		}

		// Store recommendations in cache
		err = rc.cacheService.Set(key, string(jsonRecos), rc.ttl)
		if err != nil {
			rc.logger.Err(err).Msgf("Failed to set recommendations in cache for key: %s", key)
			return []int{}, err
		}

		return apiRecos, nil
	}

	// Handle found cached recommendations
	rc.logger.Debug().Msgf("Found cached recommendations for key: %s", key)

	arrayRecos := make([]int, 0)
	err = json.Unmarshal([]byte(cachedRecos), &arrayRecos)
	if err != nil {
		rc.logger.Err(err).Msgf("Failed to unmarshal cached recommendations for key: %s", key)
		return []int{}, err
	}

	return arrayRecos, nil
}
