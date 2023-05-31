package recommendationHandler

import (
	recommendationApiService "github.com/ekkinox/fx-template/app/recommendation/api"
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	recommendationService "github.com/ekkinox/fx-template/app/recommendation/service"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/labstack/echo/v4"
	"net/http"
)

// RecommendationHandler Gather recommendations.
type RecommendationHandler struct {
	recommendationClient recommendationService.RecommendationClientContract
	productApi           recommendationApiService.ProductApiContract
	logger               *fxlogger.Logger
}

// NewRecommendationHandler Creates a new RecommendationHandler.
func NewRecommendationHandler(
	recommendationClient recommendationService.RecommendationClientContract,
	productApi recommendationApiService.ProductApiContract,
	logger *fxlogger.Logger,
) *RecommendationHandler {
	return &RecommendationHandler{
		recommendationClient: recommendationClient,
		productApi:           productApi,
		logger:               logger,
	}
}

// Handle Handles the recommendation request.
func (h *RecommendationHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get recommendations (client)
		recoType := recommendationEnum.Retailer
		retailerId := 18
		recommendationTypeId := recommendationEnum.RetailerProductsYouMayLike
		recos, _ := h.recommendationClient.GetRecommendationsByEntityAndType(
			retailerId,
			recoType,
			recommendationTypeId,
		)

		// TODO:
		//  - create the reco service,
		//  - map entities using the DB,
		//  - make sure to externalize this part to be able to replace it quickly by a monolith API
		list, _ := h.productApi.GetMany([]int{9999})
		h.logger.Debug().Interface("products", list).Msg("products")

		return c.JSON(
			http.StatusOK,
			recommendationModel.Recommendation{
				Id:       recommendationEnum.RetailerProductsYouMayLike,
				Entities: recos,
			},
		)
	}
}
