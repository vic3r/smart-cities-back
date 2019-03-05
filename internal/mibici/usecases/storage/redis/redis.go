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

// GetNeighborhoodByID returns a neighborhood
func (c *redisClient) GetNeighborhoodByID(zoneName string, neighborhoodID int64) (*models.Neighborhood, error) {

	return getNeighborhoodByID(c, zoneName, neighborhoodID)
}

// GetListNeighborhoods returns a neighborhood list
func (c *redisClient) GetListNeighborhoods() ([]*models.Neighborhood, error) {

	return getNeighborhoods(c)
}

// GetNeighborhoodsListByZone returns a list of neighborhoods per zone
func (c *redisClient) GetNeighborhoodsListByZone(zoneName string) ([]*models.Neighborhood, error) {

	return getNeighborhoodsByZone(c, zoneName)
}

// GetListZones returns a list of zones
func (c *redisClient) GetListZones() ([]*models.Zone, error) {

	return getZones(c)
}
