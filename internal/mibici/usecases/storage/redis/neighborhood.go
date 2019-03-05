package redis

import (
	"encoding/json"
	"fmt"

	"github.com/vic3r/smart-cities-back/internal/models"
)

func getNeighborhoodByID(r *redisClient, zoneName string, neighborhoodID int64) (*models.Neighborhood, error) {
	// Query to retrieve a neighborhood
	redisNeighborhood, err := r.client.LIndex(zoneName, neighborhoodID).Result()
	if err != nil {
		return nil, fmt.Errorf("neighborhood not found: %v", err)
	}

	fmt.Println(redisNeighborhood)

	nieghborhood := &models.Neighborhood{}
	// parse byte array response to json
	err = json.Unmarshal([]byte(redisNeighborhood), nieghborhood)
	if err != nil {
		return nil, fmt.Errorf("not possible to unmarshal response: %v", err)
	}

	return nieghborhood, nil
}

func getNeighborhoods(r *redisClient) ([]*models.Neighborhood, error) {
	// Query to retrieve a list of neighborhoods
	redisZones, err := r.client.Keys("*").Result()
	if err != nil {
		return nil, fmt.Errorf("zone not found: %v", err)
	}

	var neighborhoods []*models.Neighborhood
	for _, zone := range redisZones {
		neighborhoodsPerZone, err := getNeighborhoodsByZone(r, zone)
		if err != nil {
			return nil, err
		}

		neighborhoods = append(neighborhoods, neighborhoodsPerZone...)
	}

	return neighborhoods, nil
}

func getNeighborhoodsByZone(r *redisClient, zoneName string) ([]*models.Neighborhood, error) {

	zoneLength, err := r.client.LLen(zoneName).Result()
	if err != nil {
		return nil, fmt.Errorf("zone not found: %v", err)
	}

	redisNeighborhoods, err := r.client.LRange(zoneName, 0, zoneLength).Result()
	if err != nil {
		return nil, fmt.Errorf("zone range not found: %v", err)
	}

	var neighborhoods []*models.Neighborhood
	for _, actualNeighborhood := range redisNeighborhoods {
		neighborhood := &models.Neighborhood{}

		err := json.Unmarshal([]byte(actualNeighborhood), neighborhood)
		if err != nil {
			return nil, fmt.Errorf("not possible to unmarshal neighborhood: %v", err)
		}
		neighborhoods = append(neighborhoods, neighborhood)
	}

	return neighborhoods, nil
}
