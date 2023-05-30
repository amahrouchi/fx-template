package recommendationService

import (
	"github.com/ekkinox/fx-template/modules/fxconfig"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// RecommendationApi is the interface for the recommendation Api service.
type RecommendationApi interface {
	GetRecommendationsByEntityAndType(
		recommendableId int,
		recommendableType string,
		recommendationTypeId int,
		metadata map[string]any,
	) ([]int, error)
}

// NewRecommendationApi Creates a new RecommendationApi service.
func NewRecommendationApi(
	config *fxconfig.Config,
	logger *fxlogger.Logger,
	apiUrlService ApiUrl,
) RecommendationApi {
	return &DataScienceRecommendationApi{
		apiUrl:        config.GetString("config.datascience-api.url"),
		apiKey:        config.GetString("config.datascience-api.key"),
		logger:        logger,
		apiUrlService: apiUrlService,
	}
}
