package usecases_test

import (
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage"
	"github.com/vic3r/smart-cities-back/internal/models"
)

type mockStorage struct {
	GetNeighborhoodFunc       func(zoneName, neighborhoodID string) (*models.Neighborhood, error)
	GetListNeighborhoodsFunc  func() ([]*models.Neighborhood, error)
	GetListZonesFunc          func() ([]*models.Zone, error)
	GetStationsListByZoneFunc func(zoneName string) ([]*models.Station, error)
	GetStationFunc            func(zoneName string, statID int) (*models.Station, error)
}

var _ storage.Storage = &mockStorage{}

func (m *mockStorage) GetStation(zoneName string, statID int) (*models.Station, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetStationFunc(zoneName, statID)
}

func (m *mockStorage) GetNeighborhood(zoneName, neighborhoodID string) (*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetNeighborhoodFunc(zoneName, neighborhoodID)
}

func (m *mockStorage) GetListNeighborhoods() ([]*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetListNeighborhoodsFunc()
}

func (m *mockStorage) GetListZones() ([]*models.Zone, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetListZonesFunc()
}

func (m *mockStorage) GetStationsListByZone(zoneName string) ([]*models.Station, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetStationsListByZoneFunc(zoneName)
}
