package recommendationApiService

// TODO: remove this contract and its implementation when ES is fully implemented.

// LinkGeneratorContract is the interface for the LinkGenerator services.
type LinkGeneratorContract interface {
	GetProductLink(productId int, productName string, brandId int, brandName string) string
	GetBrandLink(brandId int, brandName string) string
}

// NewLinkGenerator creates a new LinkGeneratorContract service.
func NewLinkGenerator() LinkGeneratorContract {
	return &LinkGenerator{}
}
