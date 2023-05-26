package model

// Recommendation Represents a recommendation
type Recommendation struct {
	Id       int   `json:"id"`
	Entities []int `json:"entities"`
}
