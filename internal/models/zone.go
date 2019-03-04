package models

// Zone is a collection of neighborhoods
type Zone struct {
	ID            int32           `json:"id"`
	Name          string          `json:"name"`
	Neighborhoods []*Neighborhood `json:"neighborhoods"`
}
