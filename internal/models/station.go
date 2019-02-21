package models

import "time"

// Station defines a bike station with its current properties
type Station struct {
	ID             int32     `json:"id"`
	NeighborhoodID int32     `json:"neighborhood"`
	BikesQuantity  int32     `json:"bikes_quantity"`
	Date           time.Time `json:"date"`
}
