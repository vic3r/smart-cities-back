package usecases_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/vic3r/smart-cities-back/internal/mibici/usecases"
	"github.com/vic3r/smart-cities-back/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage"
)

func TestNew(t *testing.T) {
	tests := []struct {
		Name          string
		Storage       storage.Storage
		ExpectedError error
	}{
		{
			Name:          "HappyPath",
			Storage:       &mockStorage{},
			ExpectedError: nil,
		},
		{
			Name:          "NilStorage",
			Storage:       nil,
			ExpectedError: errors.New("nil storage"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			_, err := usecases.New(test.Storage)
			if test.ExpectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.ExpectedError.Error())
			}
		})
	}
}

func TestGetNeighborhood(t *testing.T) {
	tests := []struct {
		Name                 string
		ZoneName             string
		NeighborhoodID       string
		Storage              storage.Storage
		ExpectedNeighborhood *models.Neighborhood
		ExpectedError        error
	}{
		{
			Name:           "HappyPath",
			NeighborhoodID: "11",
			Storage: &mockStorage{
				GetNeighborhoodByIDFunc: func(string, int64) (*models.Neighborhood, error) {
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
						},
					}, nil
				},
			},
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
				},
			},
			ExpectedError: nil,
		},
		{
			Name:           "InvalidNeighborhoodID",
			NeighborhoodID: "a",
			Storage: &mockStorage{
				GetNeighborhoodByIDFunc: func(string, int64) (*models.Neighborhood, error) {
					return nil, errors.New("not possible to parse neighborhood_id")
				},
			},
			ExpectedNeighborhood: nil,
			ExpectedError:        errors.New("not possible to parse neighborhood_id"),
		},
		{
			Name:           "UnexpectedError",
			NeighborhoodID: "11",
			Storage: &mockStorage{
				GetNeighborhoodByIDFunc: func(string, int64) (*models.Neighborhood, error) {
					return nil, errors.New("server error")
				},
			},
			ExpectedNeighborhood: nil,
			ExpectedError:        fmt.Errorf("unexpected error server error"),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			useCases, err := usecases.New(test.Storage)
			if err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.ExpectedError.Error())
			}

			neighborhood, err := useCases.GetNeighborhood(test.ZoneName, test.NeighborhoodID)
			if err == nil {
				assert.Equal(t, test.ExpectedNeighborhood, neighborhood)
			} else {
				assert.EqualError(t, err, test.ExpectedError.Error())
			}
		})
	}
}

func TestGetNeighborhoods(t *testing.T) {
	tests := []struct {
		Name                  string
		Storage               storage.Storage
		ExpectedNeighborhoods []*models.Neighborhood
		ExpectedError         error
	}{
		{
			Name: "HappyPath",
			Storage: &mockStorage{
				GetListNeighborhoodsFunc: func() ([]*models.Neighborhood, error) {
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
							},
						},
					}, nil
				},
			},
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
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Name: "UnexpectedError",
			Storage: &mockStorage{
				GetListNeighborhoodsFunc: func() ([]*models.Neighborhood, error) {
					return nil, errors.New("server error")
				},
			},
			ExpectedNeighborhoods: nil,
			ExpectedError:         fmt.Errorf("unexpected error server error"),
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			useCases, err := usecases.New(test.Storage)
			if err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.ExpectedError.Error())
			}

			neighborhoods, err := useCases.GetNeighborhoods()
			if err == nil {
				assert.Equal(t, test.ExpectedNeighborhoods, neighborhoods)
			} else {
				assert.EqualError(t, err, test.ExpectedError.Error())
			}
		})
	}
}
