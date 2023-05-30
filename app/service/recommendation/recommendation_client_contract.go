package recommendationService

import (
	"github.com/ekkinox/fx-template/app/service/cache"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// RecommendationClientContract Represents a recommendation client service.
type RecommendationClientContract interface {
	GetRecommendationsByEntityAndType(recommendableId int, recommendableType string, recommendationTypeId int) ([]int, error)
}

// NewRecommendationClient Creates a new RecommendationClientContract service.
func NewRecommendationClient(
	recommendationApi RecommendationApiContract,
	cacheService cache.CacheContract,
	logger *fxlogger.Logger,
) RecommendationClientContract {
	return &RecommendationClient{
		recommendationApi: recommendationApi,
		cacheService:      cacheService,
		logger:            logger,
	}
}
