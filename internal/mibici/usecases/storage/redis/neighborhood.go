package redis

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vic3r/smart-cities-back/internal/models"
)

func getNeighborhood(r *redisClient, zoneName, neighborhood string) (*models.Neighborhood, error) {
	// Query to retrieve a neighborhood
	zoneLen, err := r.client.LLen(zoneName).Result()
	if err != nil {
		return nil, fmt.Errorf("zone not found: %v", err)
	}

	// parse byte array response to json
	stations := make([]*models.Station, 0)
	for i := 0; i < int(zoneLen); i++ {
		actualStation, err := r.client.LIndex(zoneName, int64(i)).Result()
		if err != nil {
			return nil, fmt.Errorf("station not found: %v", err)
		}
		if strings.Contains(actualStation, neighborhood) {
			station := &models.Station{}
			if err = json.NewDecoder(strings.NewReader(actualStation)).Decode(station); err != nil {
				return nil, fmt.Errorf("failed to decode station: %v", err)
			}
			if station.Neighborhood == neighborhood {
				stations = append(stations, station)
			}
		}
	}

	if stations == nil {
		return nil, fmt.Errorf("neighborhood not found")
	}
	return &models.Neighborhood{
		Name:         neighborhood,
		Municipality: stations[0].Zone,
		Stations:     stations,
	}, nil
}

func getNeighborhoods(r *redisClient) ([]*models.Neighborhood, error) {
	zones, err := r.client.Keys("*").Result()
	if err != nil {
		return nil, fmt.Errorf("zone not found: %v", err)
	}

	stationsDict := make(map[string][]*models.Station)
	for i, zone := range zones {
		actualStation, err := r.client.LIndex(zone, int64(i)).Result()
		if err != nil {
			return nil, fmt.Errorf("neighborhood not found: %v", err)
		}
		station := &models.Station{}
		if err = json.NewDecoder(strings.NewReader(actualStation)).Decode(station); err != nil {
			return nil, fmt.Errorf("failed to decode station: %v", err)
		}

		stationsDict[station.Neighborhood] = append(stationsDict[station.Neighborhood], station)
	}

	neighborhoods := make([]*models.Neighborhood, 0)
	for neighborhoodName, stations := range stationsDict {
		neighborhood := &models.Neighborhood{
			Name:         neighborhoodName,
			Municipality: stations[0].Zone,
			Stations:     stations,
		}
		neighborhoods = append(neighborhoods, neighborhood)
	}

	return neighborhoods, nil
}
