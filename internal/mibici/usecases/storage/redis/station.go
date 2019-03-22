package redis

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vic3r/smart-cities-back/internal/models"
)

func getStation(r *redisClient, zoneName string, stationID int) (*models.Station, error) {
	redisStation, err := r.client.LIndex(zoneName, int64(stationID)).Result()
	if err != nil {
		return nil, fmt.Errorf("station  not found: %v", err)
	}

	station := &models.Station{}
	if err = json.NewDecoder(strings.NewReader(redisStation)).Decode(station); err != nil {
		return nil, fmt.Errorf("failed decoding struct: %v", err)
	}

	return station, nil
}

func getStationsByZone(r *redisClient, zoneName string) ([]*models.Station, error) {

	zoneLength, err := r.client.LLen(zoneName).Result()
	if err != nil {
		return nil, fmt.Errorf("zone not found: %v", err)
	}

	redisStations, err := r.client.LRange(zoneName, 0, zoneLength).Result()
	if err != nil {
		return nil, fmt.Errorf("zone range not found: %v", err)
	}

	var stations []*models.Station
	for _, currentStation := range redisStations {
		station := &models.Station{}

		err := json.Unmarshal([]byte(currentStation), station)
		if err != nil {
			return nil, fmt.Errorf("not possible to unmarshal station: %v", err)
		}
		stations = append(stations, station)
	}

	return stations, nil
}
