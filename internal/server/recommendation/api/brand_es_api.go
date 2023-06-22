package recommendationApiService

import (
	recommendationEnum "github.com/ekkinox/fx-template/internal/server/recommendation/enum"
	recommendationModel "github.com/ekkinox/fx-template/internal/server/recommendation/model"
	elasticService "github.com/ekkinox/fx-template/internal/server/service/elastic"
	elasticServiceEnum "github.com/ekkinox/fx-template/internal/server/service/elastic/enum"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/samber/lo"
	"strconv"
)

// BrandEsApi service getting brand data from the database.
type BrandEsApi struct {
	esClient elasticService.ESClientContract
	logger   *fxlogger.Logger
}

// GetMany gets many brands by their ids from ElasticSearch.
func (b *BrandEsApi) GetMany(brandIds []int) ([]*recommendationModel.RecommendationBrand, error) {
	// Convert brandIds to strings
	brandIdsStr := lo.Map(brandIds, func(id int, _ int) string {
		return strconv.Itoa(id)
	})

	// Query ElasticSearch
	results, err := b.esClient.Mget(elasticServiceEnum.BrandIndexType, brandIdsStr)
	if err != nil {
		return nil, err
	}

	documents := results.(map[string]any)["docs"].([]any)

	return b.mapDocuments(documents), nil
}

// mapDocuments maps ElasticSearch documents to RecommendationBrand.
func (b *BrandEsApi) mapDocuments(documents []any) []*recommendationModel.RecommendationBrand {
	brands := make([]*recommendationModel.RecommendationBrand, 0)
	for _, document := range documents {
		// Check if the document is valid
		if document.(map[string]any)["_source"] == nil {
			b.logger.Warn().
				Interface("document", document).
				Msg("The brand document cannot be retrieved.")
			continue
		}
		doc := document.(map[string]any)["_source"].(map[string]any)

		// Get brand information
		id, okId := doc["id"].(float64)
		name, okName := doc["name"].(string)
		link, okLink := doc["link"].(string)

		// Get brand images
		rawImages, okImages := doc["images"].(map[string]any)
		images := &recommendationModel.RecommendationBrandImage{}
		if okImages {
			images.Squared = rawImages["squared"].(string)
			images.Rounded = rawImages["rounded"].(string)
			images.Large = rawImages["large"].(string)
		}

		// Check if all information are available
		if !okId || !okName || !okLink || !okImages {
			b.logger.Warn().
				Interface("document", document).
				Msg("Failed to map some brand information from ES response.")
		}

		// Append the brand to the result
		brands = append(brands, &recommendationModel.RecommendationBrand{
			Type:   recommendationEnum.BrandRecommendation,
			Id:     int(id),
			Name:   name,
			Images: images,
			Link:   link,
		})
	}

	return brands
}
