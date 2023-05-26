package recommendations

import (
	"github.com/ekkinox/fx-template/app/enum"
	"github.com/ekkinox/fx-template/app/model"
	"github.com/ekkinox/fx-template/app/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

// RecommendationHandler Gather recommendations.
type RecommendationHandler struct {
	datascienceRecommendationApi *service.DataScienceRecommendationApi
}

// NewRecommendationHandler Creates a new RecommendationHandler.
func NewRecommendationHandler(datascienceRecommendationApi *service.DataScienceRecommendationApi) *RecommendationHandler {
	return &RecommendationHandler{
		datascienceRecommendationApi: datascienceRecommendationApi,
	}
}

// Handle Handles the recommendation request.
func (h *RecommendationHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get recommendations
		recos, err := h.datascienceRecommendationApi.GetRecommendationsByEntityAndType(
			18,
			enum.Retailer,
			enum.RetailerProductsYouMayLike,
			map[string]any{},
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "An error occurred. Recommendations could not be retrieved.",
			})
		}

		return c.JSON(
			http.StatusOK,
			model.Recommendation{
				Id:       9999,
				Entities: recos,
			},
		)
	}
}
