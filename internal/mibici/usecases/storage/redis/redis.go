package redis

import (
	r "github.com/go-redis/redis"
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage"
	"github.com/vic3r/smart-cities-back/internal/models"
)

// Redis is a struct where the client lies
type redis struct {
	client *r.Client
}

var _ storage.Storage = &redis{}

// New creates new instance from a redis client
func New(client *r.Client) (storage.Storage, error) {

	return &redis{client}, nil
}

// GetNeighborhoodByID returns a neighborhood
func (r *redis) GetNeighborhoodByID(int) (*models.Neighborhood, error) {
	// TODO: Handle entitity in redis
	return nil, nil
}

// GetListNeighborhoods returns a neighborhood list
func (r *redis) GetListNeighborhoods() ([]*models.Neighborhood, error) {
	// TODO: Handle entities in redis
	return nil, nil
}
