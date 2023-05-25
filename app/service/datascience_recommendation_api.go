package service

import (
	"github.com/ekkinox/fx-template/modules/fxconfig"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// DataScienceRecommendationApi Handles the communication with the data science recommendation API.
type DataScienceRecommendationApi struct {
	config *fxconfig.Config
	logger *fxlogger.Logger
}

// NewDataScienceRecommendationApi Creates a new DataScienceRecommendationApi.
func NewDataScienceRecommendationApi(config *fxconfig.Config, logger *fxlogger.Logger) *DataScienceRecommendationApi {
	return &DataScienceRecommendationApi{
		config: config,
		logger: logger,
	}
}

// GetRecommendationsByEntityAndType Gets recommendations by entity and type.
func (s *DataScienceRecommendationApi) GetRecommendationsByEntityAndType() {
	s.logger.Info("Call to get recommendations by entity and type")
}
