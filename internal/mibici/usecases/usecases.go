package usecases

import (
	"errors"
	"fmt"
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

func (u *defaultUseCases) GetStation(zoneName, stationID string) (*models.Station, error) {
	statID, err := strconv.Atoi(stationID)
	if err != nil {
		return nil, errors.New("not possible to parse station_id")
	}
	station, err := u.storage.GetStation(zoneName, statID)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %v", err)
	}
	fmt.Println(station)

	return station, nil
}

// GetNeighborhood returns a neighborhood
func (u *defaultUseCases) GetNeighborhood(zoneName, neighborhoodID string) (*models.Neighborhood, error) {
	neighborhood, err := u.storage.GetNeighborhood(zoneName, neighborhoodID)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %v", err)
	}

	return neighborhood, nil
}

// GetNeighborhoods returns a list of neighborhoods
func (u *defaultUseCases) GetNeighborhoods() ([]*models.Neighborhood, error) {
	neighborhood, err := u.storage.GetListNeighborhoods()
	if err != nil {
		return nil, fmt.Errorf("unexpected error %v", err)
	}

	return neighborhood, nil
}

//GetZones returns a list of zones within a city
func (u *defaultUseCases) GetStationsByZone(zoneName string) ([]*models.Station, error) {
	stations, err := u.storage.GetStationsListByZone(zoneName)
	if err != nil {
		return nil, fmt.Errorf("unexpected error %v", err)
	}

	return stations, nil
}

func (u *defaultUseCases) GetZones() ([]*models.Zone, error) {
	zones, err := u.storage.GetListZones()
	if err != nil {
		return nil, fmt.Errorf("unexpected error %v", err)
	}

	return zones, nil
}
