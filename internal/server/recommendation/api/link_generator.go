package recommendationApiService

import (
	"github.com/gosimple/slug"
	"strconv"
)

// LinkGenerator is the service for generating entity links.
type LinkGenerator struct {
}

// GetProductLink generates a product link.
func (lg *LinkGenerator) GetProductLink(productId int, productName string, brandId int, brandName string) string {
	return lg.GetBrandLink(brandId, brandName) + "/" + slug.Make(productName+"-"+strconv.Itoa(productId))
}

// GetBrandLink generates a brand link.
func (lg *LinkGenerator) GetBrandLink(brandId int, brandName string) string {
	return "/brand/" + slug.Make(brandName+"-"+strconv.Itoa(brandId))
}
