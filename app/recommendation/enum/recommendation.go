package recommendationEnum

// TODO: test all types

// Recommendation types
const (
	RetailerBrandsYouMayLike           = 1
	RetailerProductsYouMayLike         = 2
	RetailerBrandsCloseToYourArea      = 3
	RetailerCategoryProductsYouMayLike = 10
	ProductBrands                      = 4
	ProductProducts                    = 5
	BrandBrands                        = 6
	BrandProducts                      = 7
	ProductProductsSameBrand           = 8
	ProductProductsCrossBrand          = 9
)

const (
	// Recommendable types
	Retailer = "retailer"
	Product  = "product"
	Brand    = "brand"

	// Recommendation types
	ProductRecommendation = "product"
	BrandRecommendation   = "brand"
)
