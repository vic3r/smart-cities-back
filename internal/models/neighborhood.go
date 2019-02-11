package models

// Neighborhood is a cluster defintion in a region (city)
type Neighborhood struct {
	ID        int32
	Name      string
	Latitude  float32
	Longitude float32
	Stations  []*Station
}
