package recommendationService

// RecommendationClientContract Represents a recommendation client service.
type RecommendationClientContract interface {
	GetRecommendationsByEntityAndType(recommendableId int, recommendableType string, recommendationTypeId int) ([]int, error)
}

// NewRecommendationClient Creates a new RecommendationClientContract service.
func NewRecommendationClient(recommendationApi RecommendationApiContract) RecommendationClientContract {
	return &RecommendationClient{
		RecommendationApi: recommendationApi,
	}
}
