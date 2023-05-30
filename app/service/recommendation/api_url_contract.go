package recommendationService

import (
	"github.com/ekkinox/fx-template/modules/fxconfig"
)

// ApiUrlContract is the interface for the ApiUrl services.
type ApiUrlContract interface {
	Url(recommendableType string, recommendationTypeId int) (string, error)
}

// NewApiUrl Creates a new ApiUrlContract service.
func NewApiUrl(config *fxconfig.Config) ApiUrlContract {
	return &DatascienceApiUrl{
		apiUrl: config.GetString("config.datascience-api.url"),
	}
}
