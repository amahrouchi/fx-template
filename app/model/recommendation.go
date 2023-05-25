package model

// Recommendation Represents a recommendation
type Recommendation struct {
	Type string `json:"type" form:"type"`
	Ids  []int  `json:"ids" form:"ids"`
}
