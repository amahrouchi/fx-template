package recommendationService

import (
	recommendationApiService "github.com/ekkinox/fx-template/app/recommendation/api"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// RecommendationServiceContract Represents a recommendation service.
type RecommendationServiceContract interface {
	GetRecommendationByTypes(
		recommendableId int,
		recommendableType string,
		typeIds []int,
	) ([]any, error)
}

// NewRecommendationService Creates a new RecommendationService service.
func NewRecommendationService(
	recommendationClient RecommendationClientContract,
	productApi recommendationApiService.ProductApiContract,
	logger *fxlogger.Logger,
) RecommendationServiceContract {
	return &RecommendationService{
		recommendationClient: recommendationClient,
		productApi:           productApi,
		logger:               logger,
	}
}
