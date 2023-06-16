package recommendationService

import (
	"encoding/json"
	"fmt"
	recommendationApiService "github.com/ekkinox/fx-template/internal/server/recommendation/api"
	recommendationEnum "github.com/ekkinox/fx-template/internal/server/recommendation/enum"
	recommendationModel "github.com/ekkinox/fx-template/internal/server/recommendation/model"
	cacheService "github.com/ekkinox/fx-template/internal/server/service/cache"
	"github.com/ekkinox/fx-template/modules/fxlogger"
)

// recommendationHolder is a struct holding recommendations and errors.
type recommendationHolder struct {
	typeId int
	recos  any
	err    error
}

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
) []any {
	// Get recommendations from client.
	recoChannel := make(chan *recommendationHolder)
	for _, typeId := range typeIds {
		go r.sendRecosToChan(recoChannel, recommendableType, recommendableId, typeId, lang)
	}

	// Wait for all goroutines gathering recommendations to finish
	var recoByTypes []any
	for recoHolder := range recoChannel {
		if recoHolder.err != nil {
			r.logger.Err(recoHolder.err).
				Str("recommendableType", recommendableType).
				Int("recommendableId", recommendableId).
				Int("recommendationTypeId", recoHolder.typeId).
				Str("lang", lang).
				Msg("Unable to get recommendations.")

			// Send empty recommendations on error
			recoByTypes = append(recoByTypes, map[string]any{
				"id":       recoHolder.typeId,
				"entities": recoHolder.recos,
				"error":    true,
			})
		} else {
			recoByTypes = append(recoByTypes, map[string]any{
				"id":       recoHolder.typeId,
				"entities": recoHolder.recos,
			})
		}

		if len(recoByTypes) == len(typeIds) {
			close(recoChannel)
		}
	}

	return recoByTypes
}

// sendRecosToChan Sends cached or API recommendations to a channel.
func (r *RecommendationService) sendRecosToChan(
	recoChannel chan *recommendationHolder,
	recommendableType string,
	recommendableId int,
	typeId int,
	lang string,
) {
	// Get cached recommendations
	key := fmt.Sprintf("recommendations[%s][%d][%d][%s]", recommendableType, recommendableId, typeId, lang)
	cachedRecos, err := r.cacheService.Get(key)
	if err != nil {
		recoChannel <- &recommendationHolder{typeId, nil, err}
		return
	}

	// Get products from api
	switch typeId {
	// Product recommendations
	case recommendationEnum.RetailerProductsYouMayLike,
		recommendationEnum.ProductProducts,
		recommendationEnum.BrandProducts,
		recommendationEnum.ProductProductsCrossBrand,
		recommendationEnum.ProductProductsSameBrand,
		recommendationEnum.RetailerCategoryProductsYouMayLike:
		// Unmarshal and return cached recommendations
		if cachedRecos != "" {
			var unmarshalledRecos []*recommendationModel.RecommendationProduct
			err = json.Unmarshal([]byte(cachedRecos), &unmarshalledRecos)
			if err != nil {
				recoChannel <- &recommendationHolder{typeId, nil, err}
				return
			}

			// Send cached product recommendations to channel
			recoChannel <- &recommendationHolder{typeId, unmarshalledRecos, nil}
			return
		}

		// Get recommendations from api
		recommendationIds, err := r.recommendationApi.GetRecommendationsByEntityAndType(
			recommendableId,
			recommendableType,
			typeId,
			map[string]any{},
		)
		if err != nil {
			recoChannel <- &recommendationHolder{typeId, nil, err}
			return
		}

		// Get complete product data from api
		products, err := r.productApi.GetMany(recommendationIds, lang)
		if err != nil {
			recoChannel <- &recommendationHolder{typeId, nil, err}
			return
		}

		// Cache the complete product recommendations
		jsonRecos, err := json.Marshal(products)
		if err != nil {
			recoChannel <- &recommendationHolder{typeId, nil, err}
			return
		}

		err = r.cacheService.Set(key, string(jsonRecos), r.ttl)
		if err != nil {
			r.logger.Err(err).
				Str("recommendableType", recommendableType).
				Int("recommendableId", recommendableId).
				Int("recommendationTypeId", typeId).
				Str("lang", lang).
				Msg("Unable to cache product recommendation.")
			// Not sending an error to the channel here
			// because we don't want to break the flow if the cache fails
		}

		// Send product recos to the channel
		recoChannel <- &recommendationHolder{typeId, products, nil}

	// Brand recommendations
	case recommendationEnum.RetailerBrandsYouMayLike,
		recommendationEnum.RetailerBrandsCloseToYourArea,
		recommendationEnum.ProductBrands,
		recommendationEnum.BrandBrands:
		// Unmarshal and return cached recommendations
		if cachedRecos != "" {
			var unmarshalledRecos []*recommendationModel.RecommendationBrand
			err = json.Unmarshal([]byte(cachedRecos), &unmarshalledRecos)
			if err != nil {
				recoChannel <- &recommendationHolder{typeId, nil, err}
				return
			}

			// Send cached brand recos to the channel
			recoChannel <- &recommendationHolder{typeId, unmarshalledRecos, nil}
			return
		}

		// Get recommendations from api
		recommendationIds, err := r.recommendationApi.GetRecommendationsByEntityAndType(recommendableId, recommendableType, typeId, map[string]any{})
		if err != nil {
			recoChannel <- &recommendationHolder{typeId, nil, err}
			return
		}

		// Get complete product data from api
		brands, err := r.brandApi.GetMany(recommendationIds)
		if err != nil {
			recoChannel <- &recommendationHolder{typeId, nil, err}
			return
		}

		// Cache the complete product recommendations
		jsonRecos, err := json.Marshal(brands)
		if err != nil {
			recoChannel <- &recommendationHolder{typeId, nil, err}
			return
		}

		err = r.cacheService.Set(key, string(jsonRecos), r.ttl)
		if err != nil {
			r.logger.Err(err).
				Str("recommendableType", recommendableType).
				Int("recommendableId", recommendableId).
				Int("recommendationTypeId", typeId).
				Str("lang", lang).
				Msg("Unable to cache product recommendation.")
			// Not sending the error to the channel here
			// because we don't want to break the flow if the cache fails
		}

		// Send brand recos to the channel
		recoChannel <- &recommendationHolder{typeId, brands, nil}
		return

	default:
		message := fmt.Sprintf("Unhandled recommendation type %d", typeId)
		r.logger.Warn().Int("typeId", typeId).Msg(message)

		// Send empty reco to the channel
		recoChannel <- &recommendationHolder{typeId, []any{}, nil}
		return
	}
}
