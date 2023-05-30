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
	recommendationApi    recommendationService.RecommendationApiContract
	recommendationClient recommendationService.RecommendationClientContract
}

// NewRecommendationHandler Creates a new RecommendationHandler.
func NewRecommendationHandler(
	recommendationApi recommendationService.RecommendationApiContract,
	recommendationClient recommendationService.RecommendationClientContract,
) *RecommendationHandler {
	return &RecommendationHandler{
		recommendationApi:    recommendationApi,
		recommendationClient: recommendationClient,
	}
}

// Handle Handles the recommendation request.
func (h *RecommendationHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get recommendations
		recos, err := h.recommendationApi.GetRecommendationsByEntityAndType(
			18,
			enum.Retailer,
			enum.RetailerProductsYouMayLike, // TODO: test RetailerCategoryProductsYouMayLike with category_id metadata
			map[string]any{},
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "An error occurred. Recommendations could not be retrieved.",
			})
		}

		// TODO: implement the client service (+ redis?)

		return c.JSON(
			http.StatusOK,
			model.Recommendation{
				Id:       9999,
				Entities: recos,
			},
		)
	}
}
