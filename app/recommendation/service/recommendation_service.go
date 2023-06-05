package recommendationService

import (
	"encoding/json"
	"fmt"
	recommendationApiService "github.com/ekkinox/fx-template/app/recommendation/api"
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	cacheService "github.com/ekkinox/fx-template/app/service/cache"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// RecommendationService service gathering recommendations from different sources.
type RecommendationService struct {
	ttl int

	recommendationApi RecommendationApiContract
	productApi        recommendationApiService.ProductApiContract
	brandApi          recommendationApiService.BrandApiContract
	cacheService      cacheService.CacheContract
	logger            *fxlogger.Logger
}

// GetRecommendationByTypes Gets recommendations by types.
func (r *RecommendationService) GetRecommendationByTypes(
	recommendableId int,
	recommendableType string,
	typeIds []int,
	lang string,
) ([]any, error) {
	// Get recommendations from client.
	var recoByTypes []any
	for _, typeId := range typeIds {
		// Get cached recommendations
		key := fmt.Sprintf("recommendations[%s][%d][%d][%s]", recommendableType, recommendableId, typeId, lang)
		cachedRecos, err := r.cacheService.Get(key)
		if err != nil {
			return nil, err
		}

		// Get products from api
		// TODO:
		//  - parallelize requests per type
		switch typeId {
		case recommendationEnum.RetailerProductsYouMayLike,
			recommendationEnum.ProductProducts,
			recommendationEnum.BrandProducts,
			recommendationEnum.ProductProductsCrossBrand,
			recommendationEnum.ProductProductsSameBrand:
			// Unmarshal and return cached recommendations
			if cachedRecos != "" {
				var unmarshalledRecos []*recommendationModel.RecommendationProduct
				err = json.Unmarshal([]byte(cachedRecos), &unmarshalledRecos)
				if err != nil {
					return nil, err
				}

				recoByTypes = append(recoByTypes, map[string]any{
					"id":       typeId,
					"entities": unmarshalledRecos,
				})
				continue
			}

			// Get recommendations from api
			recommendationIds, err := r.recommendationApi.GetRecommendationsByEntityAndType(
				recommendableId,
				recommendableType,
				typeId,
				map[string]any{},
			)
			if err != nil {
				return nil, err
			}

			// Get complete product data from api
			products, err := r.productApi.GetMany(recommendationIds, lang)
			if err != nil {
				return nil, err
			}

			// Cache the complete product recommendations
			jsonRecos, err := json.Marshal(products)
			if err != nil {
				return nil, err
			}

			err = r.cacheService.Set(key, string(jsonRecos), r.ttl)
			if err != nil {
				r.logger.Err(err).
					Str("recommendableType", recommendableType).
					Int("recommendableId", recommendableId).
					Int("recommendationTypeId", typeId).
					Str("lang", lang).
					Msg("Unable to cache product recommendation.")
				// Not returning an error here because we don't want to break the flow if the cache fails
			}

			// Add the products to the return array
			recoByTypes = append(recoByTypes, map[string]any{
				"id":       typeId,
				"entities": products,
			})

		case recommendationEnum.RetailerBrandsYouMayLike,
			recommendationEnum.RetailerBrandsCloseToYourArea,
			recommendationEnum.ProductBrands,
			recommendationEnum.BrandBrands:
			// Unmarshal ann return cached recommendations
			if cachedRecos != "" {
				var unmarshalledRecos []*recommendationModel.RecommendationBrand
				err = json.Unmarshal([]byte(cachedRecos), &unmarshalledRecos)
				if err != nil {
					return nil, err
				}

				recoByTypes = append(recoByTypes, map[string]any{
					"id":       typeId,
					"entities": unmarshalledRecos,
				})
				continue
			}

			// Get recommendations from api
			recommendationIds, err := r.recommendationApi.GetRecommendationsByEntityAndType(recommendableId, recommendableType, typeId, map[string]any{})
			if err != nil {
				return []any{}, err
			}

			// Get complete product data from api
			brands, err := r.brandApi.GetMany(recommendationIds)
			if err != nil {
				return nil, err
			}

			// Cache the complete product recommendations
			jsonRecos, err := json.Marshal(brands)
			if err != nil {
				return nil, err
			}

			err = r.cacheService.Set(key, string(jsonRecos), r.ttl)
			if err != nil {
				r.logger.Err(err).
					Str("recommendableType", recommendableType).
					Int("recommendableId", recommendableId).
					Int("recommendationTypeId", typeId).
					Str("lang", lang).
					Msg("Unable to cache product recommendation.")
				// Not returning an error here because we don't want to break the flow if the cache fails
			}

			// Add the products to the return array
			recoByTypes = append(recoByTypes, map[string]any{
				"id":       typeId,
				"entities": brands,
			})

		default:
			message := fmt.Sprintf("Unhandled recommendation type %d", typeId)
			r.logger.Warn().Int("typeId", typeId).Msg(message)
			recoByTypes = append(recoByTypes, map[string]any{
				"id":       typeId,
				"entities": []any{},
			})
		}
	}

	return recoByTypes, nil
}
