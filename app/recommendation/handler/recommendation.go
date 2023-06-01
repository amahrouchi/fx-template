package recommendationHandler

import (
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	recommendationRequest "github.com/ekkinox/fx-template/app/recommendation/request"
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
// TODO: handle other recommendation endpoints (see monolith controller)
func (h *RetailerRecommendationHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Get retailer id from JWT token
		retailerId := 18

		// Bind request
		var req recommendationRequest.RetailerRecommendationRequest
		err := c.Bind(&req)
		if err != nil {
			h.logger.Err(err).Msg("Unable to bind recommendation types to fetch.")
			return c.JSON(
				http.StatusBadRequest,
				map[string]any{"message": "bad recommendation request."},
			)
		}

		// TODO: Validate request
		//err = c.Validate(req)
		//if err != nil {
		//	h.logger.Err(err).Msg("Unable to validate recommendation types to fetch.")
		//	return c.JSON(
		//		http.StatusBadRequest,
		//		map[string]any{"message": "bad recommendation request."},
		//	)
		//}

		// Get recommendations
		recos, err := h.recommendationService.GetRecommendationByTypes(
			retailerId,
			recommendationEnum.Retailer,
			req.Types,
			req.Lang,
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
