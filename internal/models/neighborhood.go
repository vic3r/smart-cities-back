package models

// Neighborhood is a cluster defintion in a region (city)
type Neighborhood struct {
	ID         int32
	Name       string
	Latitudes  []float32
	Longitudes []float32
	Stations   []*Station
}
