package storage

import "github.com/vic3r/smart-cities-back/internal/models"

// Storage is the interface definition for the persistence layer
type Storage interface {
	GetStation(zoneName string, stationID int) (*models.Station, error)
	GetNeighborhood(zoneName, neighborhoodID string) (*models.Neighborhood, error)
	GetListNeighborhoods() ([]*models.Neighborhood, error)
	GetListZones() ([]*models.Zone, error)
	GetStationsListByZone(zoneName string) ([]*models.Station, error)
}
