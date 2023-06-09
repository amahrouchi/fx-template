package recommendationApiService

import (
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	"gorm.io/gorm"
)

// BrandApiContract service getting product data from somewhere.
type BrandApiContract interface {
	GetMany(ids []int) ([]*recommendationModel.RecommendationBrand, error)
}

// NewBrandApi Creates a new BrandApiContract service.
func NewBrandApi(gorm *gorm.DB, link LinkGeneratorContract) BrandApiContract {
	return &BrandDbApi{
		gorm: gorm,
		link: link,
	}
}
