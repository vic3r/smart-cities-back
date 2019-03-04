package mibici_test

import (
	"github.com/vic3r/smart-cities-back/internal/mibici"
	"github.com/vic3r/smart-cities-back/internal/models"
)

type mockUseCases struct {
	GetNeighborhoodFunc  func(neighborhoodID string) (*models.Neighborhood, error)
	GetNeighborhoodsFunc func() ([]*models.Neighborhood, error)
}

var _ mibici.UseCases = &mockUseCases{}

func (m *mockUseCases) GetNeighborhood(neighborhoodID string) (*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetNeighborhoodFunc(neighborhoodID)
}

func (m *mockUseCases) GetNeighborhoods() ([]*models.Neighborhood, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetNeighborhoodsFunc()
}
