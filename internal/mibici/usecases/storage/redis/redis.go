package redis

import (
	"encoding/json"
	"fmt"

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
func (r *redisClient) GetNeighborhoodByID(neighborhoodID int) (*models.Neighborhood, error) {
	// Query to retrieve a neighborhood
	neighborhoodQuery := fmt.Sprintf("neighborhood:%d", neighborhoodID)
	response := redis.NewStringCmd(r.client.Do("GET", neighborhoodQuery))

	// parse response to byte array[]
	byteResponse, err := response.Bytes()
	if err != nil {
		return nil, fmt.Errorf("not possible to parse response: %v", err)
	}

	nieghborhoodResponse := &models.Neighborhood{}
	// parse byte array response to json
	err = json.Unmarshal(byteResponse, nieghborhoodResponse)
	if err != nil {
		return nil, fmt.Errorf("not possible to unmarshal response: %v", err)
	}

	return nieghborhoodResponse, nil
}

// GetListNeighborhoods returns a neighborhood list
func (r *redisClient) GetListNeighborhoods() ([]*models.Neighborhood, error) {
	// Query to retrieve a list of neighborhoods
	response := redis.NewStringCmd(r.client.Do("GET", "neighborhood"))

	// parse response to byte array[]
	byteResponse, err := response.Bytes()
	if err != nil {
		return nil, fmt.Errorf("not possible to parse response: %v", err)
	}
	// TODO: Fix the iteration response
	nieghborhoodsResponse := []*models.Neighborhood{}
	// parse byte array response to json
	err = json.Unmarshal(byteResponse, nieghborhoodsResponse)
	if err != nil {
		return nil, fmt.Errorf("not possible to unmarshal response: %v", err)
	}

	return nieghborhoodsResponse, nil
}
