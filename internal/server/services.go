package server

import (
	"github.com/ekkinox/fx-template/internal/repository"
	recommendationApiService "github.com/ekkinox/fx-template/internal/server/recommendation/api"
	recommendationService "github.com/ekkinox/fx-template/internal/server/recommendation/service"
	"github.com/ekkinox/fx-template/internal/server/repository"
	"github.com/ekkinox/fx-template/internal/server/service"
	cacheService "github.com/ekkinox/fx-template/internal/server/service/cache"
	"github.com/ekkinox/fx-template/modules/fxgorm"
	"github.com/ekkinox/fx-template/modules/fxhealthchecker"
	"github.com/ekkinox/fx-template/modules/fxpubsub"
	"go.uber.org/fx"
)

func RegisterServices() fx.Option {
	return fx.Provide(
		// health check probes
		fxhealthchecker.AsHealthCheckerProbe(fxgorm.NewGormProbe),
		fxhealthchecker.AsHealthCheckerProbe(fxpubsub.NewPubSubProbe),
		// repositories
		repository.NewPostRepository,
		// services
		recommendationService.NewRecommendationApi,
		recommendationService.NewApiUrl,
		recommendationService.NewRecommendationService,
		cacheService.NewCacheService,
		service.NewESClient,
		// APIs
		recommendationApiService.NewProductApi,
		recommendationApiService.NewBrandApi,
		recommendationApiService.NewLinkGenerator,
	)
}
