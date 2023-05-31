package recommendationHandler

import (
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	recommendationService "github.com/ekkinox/fx-template/app/recommendation/service"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/labstack/echo/v4"
	"net/http"
)

// RetailerRecommendationHandler Gather recommendations.
type RetailerRecommendationHandler struct {
	recommendationService recommendationService.RecommendationServiceContract
	logger                *fxlogger.Logger
}

// NewRecommendationHandler Creates a new RetailerRecommendationHandler.
func NewRecommendationHandler(
	recommendationService recommendationService.RecommendationServiceContract,
	logger *fxlogger.Logger,
) *RetailerRecommendationHandler {
	return &RetailerRecommendationHandler{
		recommendationService: recommendationService,
		logger:                logger,
	}
}

// Handle Handles the recommendation request.
func (h *RetailerRecommendationHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Get all of this from the request
		retailerId := 18
		recoType := recommendationEnum.Retailer
		recommendationTypes := []int{
			recommendationEnum.RetailerProductsYouMayLike,
			recommendationEnum.RetailerBrandsYouMayLike,
			recommendationEnum.RetailerBrandsCloseToYourArea,
		}

		// Get recommendations
		recos, err := h.recommendationService.GetRecommendationByTypes(
			retailerId,
			recoType,
			recommendationTypes,
		)
		if err != nil {
			h.logger.Err(err).Msg("Unable to get recommendations from the service.")
			return c.JSON(
				http.StatusInternalServerError,
				map[string]any{"message": "Unable to retrieve recommendations."},
			)
		}

		return c.JSON(
			http.StatusOK,
			map[string]any{
				"data": recos,
			},
		)
	}
}
