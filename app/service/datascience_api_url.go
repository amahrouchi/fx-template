package service

import (
	"errors"
	"github.com/ekkinox/fx-template/app/enum"
	"github.com/ekkinox/fx-template/modules/fxconfig"
)

// DatascienceApiUrl Generates the proper URL to call the DataScience API depending on the context.
type DatascienceApiUrl struct {
	apiUrl string
}

// NewDatascienceApiUrl Creates a new DatascienceApiUrl.
func NewDatascienceApiUrl(config *fxconfig.Config) *DatascienceApiUrl {
	return &DatascienceApiUrl{
		apiUrl: config.GetString("config.datascience-api.url"),
	}
}

// Url Gets the URL to call the DataScience API.
func (u *DatascienceApiUrl) Url(recommendableType string, recommendationTypeId int) (string, error) {
	switch recommendableType {
	case enum.Retailer:
		return u.getRetailerUrl(recommendationTypeId), nil
	case enum.Product:
		url, err := u.getProductUrl(recommendationTypeId)
		return url, err
	default:
		return "", errors.New("invalid recommendable type")
	}
}

// getRetailerUrl Gets the URL to call the DataScience API for a retailer.
func (u *DatascienceApiUrl) getRetailerUrl(recommendationTypeId int) string {
	if recommendationTypeId == enum.RetailerCategoryProductsYouMayLike {
		return u.apiUrl + "/retailers_recommendations_cp/v1/"
	}

	return u.apiUrl + "/retailers_recommendations/v5/reco_by_type"
}

// getProductUrl Gets the URL to call the DataScience API for a product.
func (u *DatascienceApiUrl) getProductUrl(recommendationTypeId int) (string, error) {
	switch recommendationTypeId {
	case enum.ProductProductsSameBrand:
		return u.apiUrl + "/product_recommendations/same_brand/", nil
	case enum.ProductProductsCrossBrand:
		return u.apiUrl + "/product_recommendations/cross_brand/", nil
	default:
		return "", errors.New("invalid recommendation type")
	}

}
