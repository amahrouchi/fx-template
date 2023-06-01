package recommendationRequest

// RetailerRecommendationRequest model used to bind request data for retailer recommendations.
type RetailerRecommendationRequest struct {
	Types []int  `query:"types" validate:"required"`
	Lang  string `query:"lang" validate:"required,alpha,lowercase,len=2"`
}
