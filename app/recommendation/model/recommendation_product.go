package recommendationModel

// RecommendationProduct model.
type RecommendationProduct struct {
	Type  string                     `json:"type"`
	Id    int                        `json:"id"`
	Name  string                     `json:"name"` // TODO: handle translations
	Brand RecommendationProductBrand `json:"brand"`
}

// RecommendationProductBrand model.
type RecommendationProductBrand struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
