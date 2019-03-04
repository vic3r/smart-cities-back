package storage

import "github.com/vic3r/smart-cities-back/internal/models"

// Storage is the interface definition for the persistence layer
type Storage interface {
	GetNeighborhoodByID(int) (*models.Neighborhood, error)
	GetListNeighborhoods() ([]*models.Neighborhood, error)
	GetListZones() ([]*models.Zone, error)
	GetZoneByID(int) (*models.Zone, error)
}
