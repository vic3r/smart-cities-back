package usecases_test

import (
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage"
	"github.com/vic3r/smart-cities-back/internal/models"
)

type mockStorage struct {
	GetNeighborhoodByIDFunc        func(zoneName string, neighborhoodID int64) (*models.Neighborhood, error)
	GetListNeighborhoodsFunc       func() ([]*models.Neighborhood, error)
	GetListZonesFunc               func() ([]*models.Zone, error)
	GetNeighborhoodsListByZoneFunc func(zoneName string) ([]*models.Neighborhood, error)
}

var _ storage.Storage = &mockStorage{}

func (m *mockStorage) GetNeighborhoodByID(zoneName string, neighborhoodID int64) (*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetNeighborhoodByIDFunc(zoneName, neighborhoodID)
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

func (m *mockStorage) GetNeighborhoodsListByZone(zoneName string) ([]*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetNeighborhoodsListByZoneFunc(zoneName)
}
