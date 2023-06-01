package recommendationEnum

// TODO: test all types

const (
	// Recommendation types
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

	// Recommendable types
	Retailer = "retailer"
	Product  = "product"
	Brand    = "brand"

	// Recommendation model types
	ProductRecommendation = "product"
	BrandRecommendation   = "brand"

	DefaultLang = "en"
)
