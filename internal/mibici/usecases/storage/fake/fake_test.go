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
		ZoneName             string
		ID                   string
		ExpectedNeighborhood *models.Neighborhood
		ExpectedError        error
	}{
		{
			Name: "HappyPath",
			ExpectedNeighborhood: &models.Neighborhood{
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

			neighborhood, err := storage.GetNeighborhood(test.Name, test.ID)
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
					Name: "Puerta de Hierro",

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
