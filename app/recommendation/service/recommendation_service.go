package recommendationService

import (
	"fmt"
	recommendationApiService "github.com/ekkinox/fx-template/app/recommendation/api"
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// RecommendationService service gathering recommendations from different sources.
type RecommendationService struct {
	recommendationClient RecommendationClientContract
	productApi           recommendationApiService.ProductApiContract
	logger               *fxlogger.Logger
}

// GetRecommendationByTypes Gets recommendations by types.
func (r *RecommendationService) GetRecommendationByTypes(
	recommendableId int,
	recommendableType string,
	typeIds []int,
) ([]any, error) {
	// Get recommendations from client.
	var recoByTypes []any
	r.logger.Debug().Interface("typeIds", typeIds).Msg("typeIds")
	for _, typeId := range typeIds {
		r.logger.Debug().Int("typeId", typeId).Msg("typeId")
		recommendationIds, err := r.recommendationClient.GetRecommendationsByEntityAndType(recommendableId, recommendableType, typeId)
		if err != nil {
			r.logger.Err(err).Msg("Unable to get recommendations from client.")
			return nil, err
		}

		// Get products from api.
		switch typeId {
		case recommendationEnum.RetailerProductsYouMayLike:
			products, err := r.productApi.GetMany(recommendationIds)
			if err != nil {
				r.logger.Err(err).Msg("Unable to get products from api.")
				return nil, err
			}
			recoByTypes = append(recoByTypes, map[string]any{
				"id":       typeId,
				"entities": products,
			})
		default:
			// TODO: Handle other types.
			message := fmt.Sprintf("Unhandled recommendation type %d", typeId)
			r.logger.Warn().Int("typeId", typeId).Msg(message)
			recoByTypes = append(recoByTypes, map[string]any{
				"id":       typeId,
				"entities": []any{},
				"message":  message,
			})
		}
	}

	return recoByTypes, nil
}
