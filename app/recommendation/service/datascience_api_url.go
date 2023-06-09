package recommendationService

import (
	"errors"
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
)

// DatascienceApiUrl Generates the proper URL to call the DataScience API depending on the context.
type DatascienceApiUrl struct {
	apiUrl string
}

// Url Gets the URL to call the DataScience API.
func (u *DatascienceApiUrl) Url(recommendableType string, recommendationTypeId int) (string, error) {
	switch recommendableType {
	case recommendationEnum.Retailer:
		return u.getRetailerUrl(recommendationTypeId), nil
	case recommendationEnum.Product:
		url, err := u.getProductUrl(recommendationTypeId)
		return url, err
	default:
		return "", errors.New("invalid recommendable type")
	}
}

// getRetailerUrl Gets the URL to call the DataScience API for a retailer.
func (u *DatascienceApiUrl) getRetailerUrl(recommendationTypeId int) string {
	if recommendationTypeId == recommendationEnum.RetailerCategoryProductsYouMayLike {
		return u.apiUrl + "/retailers_recommendations_cp/v1/"
	}

	return u.apiUrl + "/retailers_recommendations/v5/reco_by_type"
}

// getProductUrl Gets the URL to call the DataScience API for a product.
func (u *DatascienceApiUrl) getProductUrl(recommendationTypeId int) (string, error) {
	switch recommendationTypeId {
	case recommendationEnum.ProductProductsSameBrand:
		return u.apiUrl + "/product_recommendations/same_brand/", nil
	case recommendationEnum.ProductProductsCrossBrand:
		return u.apiUrl + "/product_recommendations/cross_brand/", nil
	default:
		return "", errors.New("invalid recommendation type")
	}

}
