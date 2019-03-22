package fake

import (
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage"
	"github.com/vic3r/smart-cities-back/internal/models"
)

// fake is a struct where the client lies
type fake struct{}

var _ storage.Storage = &fake{}

// New creates new instance from a redis client
func New() (storage.Storage, error) {

	return &fake{}, nil
}

// GetStation returns a station
func (f *fake) GetStation(zoneName string, statID int) (*models.Station, error) {

	return &models.Station{}, nil
}

// GetNeighborhoodByID returns a neighborhood
func (f *fake) GetNeighborhood(string, string) (*models.Neighborhood, error) {
	return &models.Neighborhood{
		Name:         "Puerta de Hierro",
		Municipality: "Zapopan",
		Stations: []*models.Station{
			&models.Station{
				ID: "1",
			},
			&models.Station{
				ID: "2",
			},
			&models.Station{
				ID: "3",
			},
		},
	}, nil
}

// GetListNeighborhoods returns a neighborhood list
func (f *fake) GetListNeighborhoods() ([]*models.Neighborhood, error) {
	return []*models.Neighborhood{
		&models.Neighborhood{
			Name:         "Puerta de Hierro",
			Municipality: "Zapopan",
			Stations: []*models.Station{
				&models.Station{
					ID: "1",
				},
				&models.Station{
					ID: "2",
				},
				&models.Station{
					ID: "3",
				},
			},
		},
		&models.Neighborhood{
			Name:         "San Juan de Ocotan",
			Municipality: "Zapopan",
			Stations: []*models.Station{
				&models.Station{
					ID: "21",
				},
				&models.Station{
					ID: "22",
				},
				&models.Station{
					ID: "23",
				},
			},
		},
	}, nil
}

// GetListZones returns a zone list
func (f *fake) GetListZones() ([]*models.Zone, error) {
	return []*models.Zone{
		&models.Zone{
			ID:   1332,
			Name: "zapopan",
			Neighborhoods: []*models.Neighborhood{
				&models.Neighborhood{
					Name:         "Puerta de Hierro",
					Municipality: "Zapopan",
					Stations: []*models.Station{
						&models.Station{
							ID: "1",
						},
						&models.Station{
							ID: "2",
						},
						&models.Station{
							ID: "3",
						},
					},
				},
				&models.Neighborhood{
					Name:         "San Juan de Ocotan",
					Municipality: "Zapopan",
					Stations: []*models.Station{
						&models.Station{
							ID: "21",
						},
						&models.Station{
							ID: "22",
						},
						&models.Station{
							ID: "23",
						},
					},
				},
			},
		},
	}, nil
}

// GetZoneByID returns a particular zone
func (f *fake) GetStationsListByZone(zoneName string) ([]*models.Station, error) {
	return []*models.Station{
		&models.Station{},
	}, nil
}
