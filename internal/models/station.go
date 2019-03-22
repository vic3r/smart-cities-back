package models

// Station defines a bike station with its current properties
type Station struct {
	Zone             string `json:"zone"`
	ID               string `json:"station_id"`
	Name             string `json:"name"`
	PostalCode       string `json:"postal_code"`
	Latitude         string `json:"latitude"`
	Longitude        string `json:"longitude"`
	Neighborhood     string `json:"neighborhood"`
	Address          string `json:"address"`
	Capacity         string `json:"capacity"`
	BikesAvailable   string `json:"bikes_available"`
	BikesUnavailable string `json:"bikes_unavailable"`
	DocksAvailable   string `json:"docks_available"`
	DocksUnavailable string `json:"docks_unavailable"`
}
