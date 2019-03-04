package fake_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage/fake"
	"github.com/vic3r/smart-cities-back/internal/models"
)

func TestNew(t *testing.T) {
	tests := []struct {
		Name          string
		ExpectedError error
	}{
		{
			Name:          "HappyPath",
			ExpectedError: nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			_, err := fake.New()
			if err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.ExpectedError.Error())
			}
		})
	}
}

func TestGetNeighborhoodByID(t *testing.T) {
	tests := []struct {
		Name                 string
		ID                   int
		ExpectedNeighborhood *models.Neighborhood
		ExpectedError        error
	}{
		{
			Name: "HappyPath",
			ID:   11,
			ExpectedNeighborhood: &models.Neighborhood{
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
			ExpectedError: nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			storage, err := fake.New()
			if err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.ExpectedError.Error())
			}

			neighborhood, err := storage.GetNeighborhoodByID(test.ID)
			if err == nil {
				assert.NoError(t, err)
				assert.Equal(t, test.ExpectedNeighborhood, neighborhood)
			} else {
				assert.EqualError(t, err, test.ExpectedError.Error())
			}
		})
	}
}

func TestGetListNeighborhoods(t *testing.T) {
	tests := []struct {
		Name                  string
		ExpectedNeighborhoods []*models.Neighborhood
		ExpectedError         error
	}{
		{
			Name: "HappyPath",
			ExpectedNeighborhoods: []*models.Neighborhood{
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
			},
			ExpectedError: nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			storage, err := fake.New()
			if err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.ExpectedError.Error())
			}

			neighborhoods, err := storage.GetListNeighborhoods()
			if err == nil {
				assert.NoError(t, err)
				assert.Equal(t, test.ExpectedNeighborhoods, neighborhoods)
			} else {
				assert.EqualError(t, err, test.ExpectedError.Error())
			}
		})
	}
}
