package recommendationApiService

import (
	recommendationModel "github.com/ekkinox/fx-template/internal/server/http/recommendation/model"
	elasticService "github.com/ekkinox/fx-template/internal/server/http/service/elastic"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// BrandApiContract service getting product data from somewhere.
type BrandApiContract interface {
	GetMany(ids []int) ([]*recommendationModel.RecommendationBrand, error)
}

// NewBrandApi Creates a new BrandApiContract service.
func NewBrandApi(esClient elasticService.ESClientContract, logger *fxlogger.Logger) BrandApiContract {
	return &BrandEsApi{
		esClient: esClient,
		logger:   logger,
	}
}
