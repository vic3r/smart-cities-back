package redis

import (
	"github.com/go-redis/redis"
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage"
	"github.com/vic3r/smart-cities-back/internal/models"
)

// Redis is a struct where the client lies
type redisClient struct {
	client *redis.Client
}

var _ storage.Storage = &redisClient{}

// New creates new instance from a redisClient client
func New(client *redis.Client) (storage.Storage, error) {

	return &redisClient{client}, nil
}

// GetStation returns a station
func (c *redisClient) GetStation(zoneName string, stationID int) (*models.Station, error) {

	return getStation(c, zoneName, stationID)
}

// GetNeighborhoodByID returns a neighborhood
func (c *redisClient) GetNeighborhood(zoneName string, neighborhood string) (*models.Neighborhood, error) {

	return getNeighborhood(c, zoneName, neighborhood)
}

// GetListNeighborhoods returns a neighborhood list
func (c *redisClient) GetListNeighborhoods() ([]*models.Neighborhood, error) {

	return getNeighborhoods(c)
}

// GetStationsListByZone returns a list of neighborhoods per zone
func (c *redisClient) GetStationsListByZone(zoneName string) ([]*models.Station, error) {

	return getStationsByZone(c, zoneName)
}

// GetListZones returns a list of zones
func (c *redisClient) GetListZones() ([]*models.Zone, error) {

	return getZones(c)
}
