package models

import "time"

// Station defines a bike station with its current properties
type Station struct {
	ID             int32     `json:"id"`
	Municipality   string    `json:"municipality"`
	Neighbourhood  string    `json:"neighbourhood"`
	BikesQuanitity int32     `json:"bikes_quanitity"`
	Date           time.Time `json:"date"`
}
