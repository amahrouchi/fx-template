package recommendationApiService

import (
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"gorm.io/gorm"
)

// ProductApiContract service getting product data from somewhere.
type ProductApiContract interface {
	GetMany(ids []int) ([]recommendationModel.RecommendationProduct, error)
}

// NewProductApi Creates a new ProductApiContract service.
func NewProductApi(gorm *gorm.DB, logger *fxlogger.Logger) ProductApiContract {
	return &ProductDbApi{
		gorm:   gorm,
		logger: logger,
	}
}
