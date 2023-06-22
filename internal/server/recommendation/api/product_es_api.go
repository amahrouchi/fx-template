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

// ProductEsApi service getting product data from elasticsearch.
type ProductEsApi struct {
	esClient elasticService.ESClientContract
	link     LinkGeneratorContract
	logger   *fxlogger.Logger
}

// GetMany gets many products from the database.
func (p *ProductEsApi) GetMany(productIds []int, lang string) ([]*recommendationModel.RecommendationProduct, error) {
	// Convert productIds to strings
	productIdsStr := lo.Map(productIds, func(id int, _ int) string {
		return strconv.Itoa(id)
	})

	// Query ElasticSearch
	results, err := p.esClient.Mget(elasticServiceEnum.ProductIndexType, productIdsStr)
	if err != nil {
		return nil, err
	}

	documents := results.(map[string]any)["docs"].([]any)

	return p.mapDocuments(documents, lang), nil
}

// mapDocuments maps ElasticSearch documents to RecommendationProduct.
func (p *ProductEsApi) mapDocuments(documents []any, lang string) []*recommendationModel.RecommendationProduct {
	products := make([]*recommendationModel.RecommendationProduct, 0)
	for _, document := range documents {
		// Check if the document is valid
		if document.(map[string]any)["_source"] == nil {
			p.logger.Warn().
				Interface("document", document).
				Msg("The product document cannot be retrieved.")
			continue
		}
		doc := document.(map[string]any)["_source"].(map[string]any)

		// Get product information
		id, okId := doc["id"].(float64)
		link, okLink := doc["link"].(string)
		rawImages := doc["images"].([]any)
		images := lo.Map(rawImages, func(image any, _ int) string {
			return image.(string)
		})

		// Get product name to display
		defaultName, okDefaultName := doc["name"].(map[string]any)["en"].(string)
		name, okName := doc["name"].(map[string]any)[lang].(string)
		displayedName := lo.Ternary(okName, name, defaultName)

		// Get brand information
		brandId, okBrandId := doc["brand"].(map[string]any)["id"].(float64)
		brandName, okBrandName := doc["brand"].(map[string]any)["name"].(string)
		brandLink, okBrandLink := doc["brand"].(map[string]any)["link"].(string)

		// Check if all information are available
		if !okId || !okDefaultName || !okName || !okLink || !okBrandId || !okBrandName || !okBrandLink {
			p.logger.Warn().
				Interface("document", document).
				Msg("Failed to map some product information from ES response.")
		}

		// Append the product to the result
		products = append(products, &recommendationModel.RecommendationProduct{
			Type:   recommendationEnum.ProductRecommendation,
			Id:     int(id),
			Name:   displayedName,
			Images: images,
			Link:   link,
			Brand: &recommendationModel.RecommendationProductBrand{
				Id:   int(brandId),
				Name: brandName,
				Link: brandLink,
			},
		})
	}

	return products
}
