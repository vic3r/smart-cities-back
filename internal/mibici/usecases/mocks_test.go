package usecases_test

import (
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage"
	"github.com/vic3r/smart-cities-back/internal/models"
)

type mockStorage struct {
	GetNeighborhoodByIDFunc  func(neighborhoodID int) (*models.Neighborhood, error)
	GetListNeighborhoodsFunc func() ([]*models.Neighborhood, error)
}

var _ storage.Storage = &mockStorage{}

func (m *mockStorage) GetNeighborhoodByID(neighborhoodID int) (*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetNeighborhoodByIDFunc(neighborhoodID)
}

func (m *mockStorage) GetListNeighborhoods() ([]*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetListNeighborhoodsFunc()
}
