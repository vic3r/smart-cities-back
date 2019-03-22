package mibici

import (
	"github.com/vic3r/smart-cities-back/internal/models"
)

// UseCases is an interface for the usecases
type UseCases interface {
	GetStation(zoneName, neighborhoodID string) (*models.Station, error)
	GetNeighborhood(zoneName, neighborhoodID string) (*models.Neighborhood, error)
	GetNeighborhoods() ([]*models.Neighborhood, error)
	GetStationsByZone(zone string) ([]*models.Station, error)
	GetZones() ([]*models.Zone, error)
}
