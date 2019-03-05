package mibici_test

import (
	"github.com/vic3r/smart-cities-back/internal/mibici"
	"github.com/vic3r/smart-cities-back/internal/models"
)

type mockUseCases struct {
	GetNeighborhoodFunc        func(zoneName, neighborhoodID string) (*models.Neighborhood, error)
	GetNeighborhoodsFunc       func() ([]*models.Neighborhood, error)
	GetNeighborhoodsByZoneFunc func(zone string) ([]*models.Neighborhood, error)
	GetZonesFunc               func() ([]*models.Zone, error)
}

var _ mibici.UseCases = &mockUseCases{}

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

func (m *mockUseCases) GetNeighborhoodsByZone(zone string) ([]*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetNeighborhoodsByZone(zone)
}

func (m *mockUseCases) GetZones() ([]*models.Zone, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetZonesFunc()
}
