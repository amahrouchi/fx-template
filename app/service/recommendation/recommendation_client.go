package recommendationService

// RecommendationClient Represents a recommendation client service.
type RecommendationClient interface {
	GetRecommendationsByEntityAndType(recommendableId int, recommendableType string, recommendationTypeId int) ([]int, error)
}

// NewRecommendationClient Creates a new RecommendationClient service.
func NewRecommendationClient(recommendationApi RecommendationApi) RecommendationClient {
	return &RecommendationClientService{
		RecommendationApi: recommendationApi,
	}
}
