package recommendationModel

type RecommendationBrand struct {
	Type   string                    `json:"type"`
	Id     int                       `json:"id"`
	Name   string                    `json:"name"`
	Images *RecommendationBrandImage `json:"images"`
	Link   string                    `json:"link"` // TODO: generate brand link
}

type RecommendationBrandImage struct {
	Squared *string `json:"squared"`
	Rounded *string `json:"rounded"`
	Large   *string `json:"large"`
}
