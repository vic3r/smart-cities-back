package mibici

import (
	"github.com/vic3r/smart-cities-back/internal/models"
)

// UseCases is an interface for the usecases
type UseCases interface {
	GetNeighborhood(neighborhoodID string) (*models.Neighborhood, error)
	GetNeighborhoods() ([]*models.Neighborhood, error)
}
