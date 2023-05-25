package recommendations

import (
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
		h.datascienceRecommendationApi.GetRecommendationsByEntityAndType()

		return c.JSON(
			http.StatusOK,
			model.Recommendation{
				Type: "whatever",
				Ids:  []int{1, 2, 3},
			},
		)
	}
}
