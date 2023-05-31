package app

import (
	recommendationApiService "github.com/ekkinox/fx-template/app/recommendation/api"
	recommendationService "github.com/ekkinox/fx-template/app/recommendation/service"
	"github.com/ekkinox/fx-template/app/repository"
	"github.com/ekkinox/fx-template/app/service/cache"
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
		cacheService.NewCacheService,
		// APIs
		recommendationApiService.NewProductApi,
	)
}
