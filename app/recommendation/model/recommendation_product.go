package recommendationModel

// RecommendationProduct model.
type RecommendationProduct struct {
	Type   string                     `json:"type"`
	Id     int                        `json:"id"`
	Name   string                     `json:"name"`
	Images []string                   `json:"images"` // TODO: get product images
	Link   string                     `json:"link"`   // TODO: generate product link
	Brand  RecommendationProductBrand `json:"brand"`
}

// RecommendationProductBrand model.
type RecommendationProductBrand struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
