package recommendationService

// RecommendationClientService Gather recommendations from the recommendation API.
type RecommendationClientService struct {
	RecommendationApi RecommendationApi
}

// GetRecommendationsByEntityAndType Get recommendations by entity and type.
func (rc *RecommendationClientService) GetRecommendationsByEntityAndType(
	recommendableId int,
	recommendableType string,
	recommendationTypeId int,
) ([]int, error) {
	return []int{1, 2, 3, 4, 5}, nil
}
