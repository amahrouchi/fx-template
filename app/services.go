package app

import (
	"github.com/ekkinox/fx-template/app/repository"
	recommendationService "github.com/ekkinox/fx-template/app/service/recommendation"
	"github.com/ekkinox/fx-template/modules/fxgorm"
	"github.com/ekkinox/fx-template/modules/fxhealthchecker"
	"go.uber.org/fx"
)

func RegisterServices() fx.Option {
	return fx.Provide(
		// health check probes
		fxhealthchecker.AsProbe(fxgorm.NewGormProbe),
		// repositories
		repository.NewPostRepository,
		// services
		recommendationService.NewRecommendationApi,
		recommendationService.NewApiUrl,
		recommendationService.NewRecommendationClient,
	)
}
