package models

// Zone defines a zone model
type Zone struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Neighborhoods []*Neighborhood `json:"neighborhoods,omitempty"`
}
