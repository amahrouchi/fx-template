package recommendationService

import (
	recommendationApiService "github.com/ekkinox/fx-template/internal/server/recommendation/api"
	cacheService "github.com/ekkinox/fx-template/internal/server/service/cache"
	"github.com/ekkinox/fx-template/modules/fxconfig"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// RecommendationServiceContract Represents a recommendation service.
type RecommendationServiceContract interface {
	GetRecommendationByTypes(
		recommendableId int,
		recommendableType string,
		typeIds []int,
		lang string,
	) []any
}

// NewRecommendationService Creates a new RecommendationService service.
func NewRecommendationService(
	recommendationApi RecommendationApiContract,
	productApi recommendationApiService.ProductApiContract,
	brandApi recommendationApiService.BrandApiContract,
	cacheService cacheService.CacheContract,
	logger *fxlogger.Logger,
	config *fxconfig.Config,
) RecommendationServiceContract {
	return &RecommendationService{
		ttl:               config.GetInt("config.recommendation.ttl"),
		recommendationApi: recommendationApi,
		productApi:        productApi,
		brandApi:          brandApi,
		cacheService:      cacheService,
		logger:            logger,
	}
}
