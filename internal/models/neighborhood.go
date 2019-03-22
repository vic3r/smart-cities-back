package models

// Neighborhood is a cluster defintion in a region (city)
type Neighborhood struct {
	Name         string     `json:"name"`
	Municipality string     `json:"municipality"`
	Stations     []*Station `json:"stations"`
}
