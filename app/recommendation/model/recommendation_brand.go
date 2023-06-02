package recommendationModel

type RecommendationBrand struct {
	Type   string            `json:"type"`
	Id     int               `json:"id"`
	Name   string            `json:"name"`
	Images map[string]string `json:"images"` // TODO: get brand images
	Link   string            `json:"link"`   // TODO: generate brand link
}
