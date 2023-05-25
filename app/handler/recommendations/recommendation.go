package recommendations

import (
	"github.com/ekkinox/fx-template/app/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

// RecommendationHandler Gather recommendations.
type RecommendationHandler struct {
}

// NewRecommendationHandler Creates a new RecommendationHandler.
func NewRecommendationHandler() *RecommendationHandler {
	return &RecommendationHandler{}
}

// Handle Handles the recommendation request.
func (h *RecommendationHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			model.Recommendation{
				Type: "whatever",
				Ids:  []int{1, 2, 3},
			},
		)
	}
}
