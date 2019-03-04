package usecases

import (
	"errors"
	"strconv"

	"github.com/vic3r/smart-cities-back/internal/mibici"
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage"
	"github.com/vic3r/smart-cities-back/internal/models"
)

// defaultUseCases is a usecases implementation
type defaultUseCases struct {
	storage storage.Storage
}

var _ mibici.UseCases = &defaultUseCases{}

// New creates a new set of default use cases
func New(storage storage.Storage) (mibici.UseCases, error) {
	if storage == nil {
		return nil, errors.New("nil storage")
	}

	return &defaultUseCases{storage: storage}, nil
}

// GetNeighborhood returns a neighborhood
func (u *defaultUseCases) GetNeighborhood(neighborhoodID string) (*models.Neighborhood, error) {
	nbhoodID, err := strconv.Atoi(neighborhoodID)
	if err != nil {
		return nil, errors.New("not possible to parse neighborhood_id")
	}
	neighborhood, err := u.storage.GetNeighborhoodByID(nbhoodID)
	if err != nil {
		return nil, err
	}

	return neighborhood, nil
}

// GetNeighborhoods returns a list of neighborhoods
func (u *defaultUseCases) GetNeighborhoods() ([]*models.Neighborhood, error) {
	neighborhood, err := u.storage.GetListNeighborhoods()
	if err != nil {
		return nil, err
	}

	return neighborhood, nil
}

//GetZones returns a list of zones within a city
func (u *defaultUseCases) GetNeighborhoodsByZone(zoneID string) ([]*models.Zone, error) {
	zneID, err := strconv.Atoi(zoneID)
	if err != nil {
		return nil, errors.New("not possible to parse zone_id")
	}
	zone, err := u.storage.GetNeighborhoodByID(zneID)
	if err != nil {
		return nil, err
	}

	return zone, nil

}
