package fake

import (
	"time"

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

// GetNeighborhoodByID returns a neighborhood
func (f *fake) GetNeighborhoodByID(string, int64) (*models.Neighborhood, error) {
	return &models.Neighborhood{
		ID:           11,
		Name:         "Puerta de Hierro",
		Latitudes:    []float32{1.5, 3.0},
		Longitudes:   []float32{2.4, 1.3},
		Municipality: "Zapopan",
		Stations: []*models.Station{
			&models.Station{
				ID:             1,
				NeighborhoodID: 11,
				BikesQuantity:  8,
			},
			&models.Station{
				ID:             2,
				NeighborhoodID: 11,
				BikesQuantity:  6,
			},
			&models.Station{
				ID:             3,
				NeighborhoodID: 11,
				BikesQuantity:  10,
			},
		},
	}, nil
}

// GetListNeighborhoods returns a neighborhood list
func (f *fake) GetListNeighborhoods() ([]*models.Neighborhood, error) {
	return []*models.Neighborhood{
		&models.Neighborhood{
			ID:           11,
			Name:         "Puerta de Hierro",
			Latitudes:    []float32{1.5, 3.0},
			Longitudes:   []float32{2.4, 1.3},
			Municipality: "Zapopan",
			Stations: []*models.Station{
				&models.Station{
					ID:             1,
					NeighborhoodID: 11,
					BikesQuantity:  8,
				},
				&models.Station{
					ID:             2,
					NeighborhoodID: 11,
					BikesQuantity:  6,
				},
				&models.Station{
					ID:             3,
					NeighborhoodID: 11,
					BikesQuantity:  10,
				},
			},
		},
		&models.Neighborhood{
			ID:           12,
			Name:         "San Juan de Ocotan",
			Latitudes:    []float32{4.5, -1.0},
			Longitudes:   []float32{5.4, 4.3},
			Municipality: "Zapopan",
			Stations: []*models.Station{
				&models.Station{
					ID:             21,
					NeighborhoodID: 12,
					BikesQuantity:  9,
				},
				&models.Station{
					ID:             22,
					NeighborhoodID: 12,
					BikesQuantity:  6,
				},
				&models.Station{
					ID:             23,
					NeighborhoodID: 12,
					BikesQuantity:  8,
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
					ID:           11,
					Name:         "Puerta de Hierro",
					Latitudes:    []float32{1.5, 3.0},
					Longitudes:   []float32{2.4, 1.3},
					Municipality: "Zapopan",
					Stations: []*models.Station{
						&models.Station{
							ID:             1,
							NeighborhoodID: 11,
							BikesQuantity:  8,
							Date:           time.Now(),
						},
						&models.Station{
							ID:             2,
							NeighborhoodID: 11,
							BikesQuantity:  6,
							Date:           time.Now(),
						},
						&models.Station{
							ID:             3,
							NeighborhoodID: 11,
							BikesQuantity:  10,
							Date:           time.Now(),
						},
					},
				},
				&models.Neighborhood{
					ID:           12,
					Name:         "San Juan de Ocotan",
					Latitudes:    []float32{4.5, -1.0},
					Longitudes:   []float32{5.4, 4.3},
					Municipality: "Zapopan",
					Stations: []*models.Station{
						&models.Station{
							ID:             21,
							NeighborhoodID: 12,
							BikesQuantity:  9,
							Date:           time.Now(),
						},
						&models.Station{
							ID:             22,
							NeighborhoodID: 12,
							BikesQuantity:  6,
							Date:           time.Now(),
						},
						&models.Station{
							ID:             23,
							NeighborhoodID: 12,
							BikesQuantity:  8,
							Date:           time.Now(),
						},
					},
				},
			},
		},
	}, nil
}

// GetZoneByID returns a particular zone
func (f *fake) GetNeighborhoodsListByZone(zoneName string) ([]*models.Neighborhood, error) {
	return []*models.Neighborhood{
		&models.Neighborhood{
			ID:           11,
			Name:         "Puerta de Hierro",
			Latitudes:    []float32{1.5, 3.0},
			Longitudes:   []float32{2.4, 1.3},
			Municipality: "Zapopan",
			Stations: []*models.Station{
				&models.Station{
					ID:             1,
					NeighborhoodID: 11,
					BikesQuantity:  8,
				},
				&models.Station{
					ID:             2,
					NeighborhoodID: 11,
					BikesQuantity:  6,
				},
				&models.Station{
					ID:             3,
					NeighborhoodID: 11,
					BikesQuantity:  10,
				},
			},
		},
		&models.Neighborhood{
			ID:           12,
			Name:         "San Juan de Ocotan",
			Latitudes:    []float32{4.5, -1.0},
			Longitudes:   []float32{5.4, 4.3},
			Municipality: "Zapopan",
			Stations: []*models.Station{
				&models.Station{
					ID:             21,
					NeighborhoodID: 12,
					BikesQuantity:  9,
				},
				&models.Station{
					ID:             22,
					NeighborhoodID: 12,
					BikesQuantity:  6,
				},
				&models.Station{
					ID:             23,
					NeighborhoodID: 12,
					BikesQuantity:  8,
				},
			},
		},
	}, nil
}
