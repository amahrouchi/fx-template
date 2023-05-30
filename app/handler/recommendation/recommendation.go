package recommendationHandler

import (
	"github.com/ekkinox/fx-template/app/enum"
	"github.com/ekkinox/fx-template/app/model"
	recommendationService "github.com/ekkinox/fx-template/app/service/recommendation"
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
		recos, _ := h.recommendationClient.GetRecommendationsByEntityAndType(18, enum.Retailer, enum.RetailerProductsYouMayLike)

		return c.JSON(
			http.StatusOK,
			model.Recommendation{
				Id:       9999,
				Entities: recos,
			},
		)
	}
}
