package recommendationModel

// TODO: See with FE the list of needed information.

// RecommendationBrand the model displayed for brand recommendations.
type RecommendationBrand struct {
	Type   string                    `json:"type"`
	Id     int                       `json:"id"`
	Name   string                    `json:"name"`
	Images *RecommendationBrandImage `json:"images"`
	Link   string                    `json:"link"`
}

// RecommendationBrandImage the model displayed for brand recommendation images.
type RecommendationBrandImage struct {
	Squared string `json:"squared"`
	Rounded string `json:"rounded"`
	Large   string `json:"large"`
}
