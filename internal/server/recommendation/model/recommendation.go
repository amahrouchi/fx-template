package recommendationModel

// Recommendation Represents a recommendation
type Recommendation struct {
	Id       int                      `json:"id"`
	Entities []*RecommendationProduct `json:"entities"`
}
