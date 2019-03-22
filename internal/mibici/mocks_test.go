package mibici_test

import (
	"github.com/vic3r/smart-cities-back/internal/mibici"
	"github.com/vic3r/smart-cities-back/internal/models"
)

type mockUseCases struct {
	GetNeighborhoodFunc   func(zoneName, neighborhoodID string) (*models.Neighborhood, error)
	GetNeighborhoodsFunc  func() ([]*models.Neighborhood, error)
	GetStationsByZoneFunc func(zone string) ([]*models.Station, error)
	GetZonesFunc          func() ([]*models.Zone, error)
	GetStationFunc        func(zoneName, neighborhoodID string) (*models.Station, error)
}

var _ mibici.UseCases = &mockUseCases{}

func (m *mockUseCases) GetStation(zoneName, stationID string) (*models.Station, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetStationFunc(zoneName, stationID)
}

func (m *mockUseCases) GetNeighborhood(zoneName, neighborhoodID string) (*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetNeighborhoodFunc(zoneName, neighborhoodID)
}

func (m *mockUseCases) GetNeighborhoods() ([]*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetNeighborhoodsFunc()
}

func (m *mockUseCases) GetStationsByZone(zone string) ([]*models.Station, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetStationsByZoneFunc(zone)
}

func (m *mockUseCases) GetZones() ([]*models.Zone, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetZonesFunc()
}
