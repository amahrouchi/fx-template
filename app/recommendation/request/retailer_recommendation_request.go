package recommendationRequest

// RetailerRecommendationRequest model used to bind request data for retailer recommendations.
type RetailerRecommendationRequest struct {
	Types []int `query:"types"`
}
