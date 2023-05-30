package recommendationService

import (
	"github.com/ekkinox/fx-template/modules/fxconfig"
)

type ApiUrl interface {
	Url(recommendableType string, recommendationTypeId int) (string, error)
}

// NewApiUrl Creates a new ApiUrl service.
func NewApiUrl(config *fxconfig.Config) ApiUrl {
	return &DatascienceApiUrl{
		apiUrl: config.GetString("config.datascience-api.url"),
	}
}
