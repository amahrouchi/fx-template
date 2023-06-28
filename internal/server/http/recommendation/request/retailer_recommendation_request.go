package recommendationRequest

// RetailerRecommendationRequest model used to bind request data for retailer recommendations.
type RetailerRecommendationRequest struct {
	Types []int  `query:"types[]" validate:"required,dive,oneof=1 2 3"` // TODO: how to use the constants from recommendationEnum?
	Lang  string `query:"lang" validate:"required,alpha,lowercase,len=2"`
}
