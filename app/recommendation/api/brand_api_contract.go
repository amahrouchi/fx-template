package recommendationApiService

import (
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"gorm.io/gorm"
)

// BrandApiContract service getting product data from somewhere.
type BrandApiContract interface {
	GetMany(ids []int) ([]*recommendationModel.RecommendationBrand, error)
}

// NewBrandApi Creates a new BrandApiContract service.
func NewBrandApi(gorm *gorm.DB, link LinkGeneratorContract, logger *fxlogger.Logger) BrandApiContract {
	return &BrandDbApi{
		gorm:   gorm,
		link:   link,
		logger: logger,
	}
}
