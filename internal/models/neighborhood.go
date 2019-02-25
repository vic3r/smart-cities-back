package models

// Neighborhood is a cluster defintion in a region (city)
type Neighborhood struct {
	ID           int32      `json:"id"`
	Name         string     `json:"name"`
	Latitudes    []float32  `json:"latitudes"`
	Longitudes   []float32  `json:"longitudes"`
	Municipality string     `json:"municipality"`
	Stations     []*Station `json:"stations"`
}
