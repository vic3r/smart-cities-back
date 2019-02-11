package models

import "time"

// Station defines a bike station with its current properties
type Station struct {
	ID            int32         `json:"id"`
	Municipality  string        `json:"municipality"`
	Neighborhood  *Neighborhood `json:"neighborhood"`
	BikesQuantity int32         `json:"bikes_quantity"`
	Date          time.Time     `json:"date"`
}
