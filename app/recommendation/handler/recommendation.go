package recommendationHandler

import (
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	recommendationRequest "github.com/ekkinox/fx-template/app/recommendation/request"
	recommendationService "github.com/ekkinox/fx-template/app/recommendation/service"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
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
				map[string]any{"message": "Bad recommendation request."},
			)
		}

		// Validate request
		v := validator.New()
		err = v.Struct(req)
		if err != nil {
			h.logger.Err(err).Msg("Unable to validate recommendation types to fetch.")
			return c.JSON(
				http.StatusBadRequest,
				map[string]any{"message": "Bad recommendation request."},
			)
		}

		// Get recommendations
		recos := h.recommendationService.GetRecommendationByTypes(
			retailerId,
			recommendationEnum.Retailer,
			req.Types,
			req.Lang,
		)

		// Check if there are only errors
		onlyErrors := lo.Reduce(recos, func(agg bool, item any, _ int) bool {
			itemHasError, ok := item.(map[string]any)["error"]

			return agg && ok && itemHasError.(bool)
		}, true)

		if onlyErrors {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]any{"message": "All recommendation requests failed."},
			)
		}

		return c.JSON(
			http.StatusOK,
			map[string]any{"data": recos},
		)
	}
}
