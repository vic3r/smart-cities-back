package redis

import (
	"fmt"

	"github.com/vic3r/smart-cities-back/internal/models"
)

func getZones(r *redisClient) ([]*models.Zone, error) {
	redisZones, err := r.client.Keys("*").Result()
	if err != nil {
		return nil, fmt.Errorf("zone not found: %v", err)
	}

	var zones []*models.Zone
	for i := 0; i < len(redisZones); i++ {
		zoneID := i + 1
		zone := &models.Zone{
			ID:   zoneID,
			Name: redisZones[i],
		}
		zones = append(zones, zone)
	}

	return zones, nil
}
