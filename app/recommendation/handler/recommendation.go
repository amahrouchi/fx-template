package recommendationHandler

import (
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	recommendationService "github.com/ekkinox/fx-template/app/recommendation/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

// RecommendationHandler Gather recommendations.
type RecommendationHandler struct {
	recommendationClient recommendationService.RecommendationClientContract
}

// NewRecommendationHandler Creates a new RecommendationHandler.
func NewRecommendationHandler(
	recommendationClient recommendationService.RecommendationClientContract,
) *RecommendationHandler {
	return &RecommendationHandler{
		recommendationClient: recommendationClient,
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

		return c.JSON(
			http.StatusOK,
			recommendationModel.Recommendation{
				Id:       recommendationEnum.RetailerProductsYouMayLike,
				Entities: recos,
			},
		)
	}
}
