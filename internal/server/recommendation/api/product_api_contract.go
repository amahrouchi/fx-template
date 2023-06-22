package recommendationApiService

import (
	recommendationModel "github.com/ekkinox/fx-template/internal/server/recommendation/model"
	elasticService "github.com/ekkinox/fx-template/internal/server/service/elastic"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// ProductApiContract service getting product data from somewhere.
type ProductApiContract interface {
	GetMany(ids []int, lang string) ([]*recommendationModel.RecommendationProduct, error)
}

// NewProductApi Creates a new ProductApiContract service.
func NewProductApi(esClient elasticService.ESClientContract, logger *fxlogger.Logger) ProductApiContract {
	return &ProductEsApi{
		logger:   logger,
		esClient: esClient,
	}
}
