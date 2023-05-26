package service

import (
	"github.com/ekkinox/fx-template/modules/fxconfig"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// DataScienceRecommendationApi Handles the communication with the data science recommendation API.
type DataScienceRecommendationApi struct {
	apiUrl string
	apiKey string

	logger *fxlogger.Logger
}

// NewDataScienceRecommendationApi Creates a new DataScienceRecommendationApi.
func NewDataScienceRecommendationApi(config *fxconfig.Config, logger *fxlogger.Logger) *DataScienceRecommendationApi {
	return &DataScienceRecommendationApi{
		apiUrl: config.GetString("config.datascience-api.url"),
		apiKey: config.GetString("config.datascience-api.key"),
		logger: logger,
	}
}

// GetRecommendationsByEntityAndType Gets recommendations by entity and type.
func (s *DataScienceRecommendationApi) GetRecommendationsByEntityAndType() {
	s.logger.Debug().Msg("Call to get recommendations by entity and type")
	s.logger.Debug().Msgf("url: %v", s.apiUrl)
	s.logger.Debug().Msgf("key: %v", s.apiKey)
}
