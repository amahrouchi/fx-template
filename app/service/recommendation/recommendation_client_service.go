package recommendationService

// RecommendationClient Gather recommendations from the recommendation API.
type RecommendationClient struct {
	RecommendationApi RecommendationApiContract
}

// GetRecommendationsByEntityAndType Get recommendations by entity and type.
func (rc *RecommendationClient) GetRecommendationsByEntityAndType(
	recommendableId int,
	recommendableType string,
	recommendationTypeId int,
) ([]int, error) {
	//key := fmt.Sprintf("{%s}.{%d}.{%d}", recommendableType, recommendableId, recommendationTypeId)

	return []int{1, 2, 3, 4, 5}, nil
}
