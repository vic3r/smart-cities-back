package redis

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/mitchellh/mapstructure"
	"github.com/vic3r/smart-cities-back/services/db"
)

// redisConfig is a redis configuration
type redisConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database int
}

type service struct{}

var _ db.Service = &service{}

// Connection returns a new client connection
func (s *service) Connection(rawConfig map[string]interface{}) (*redis.Client, error) {
	config := &redisConfig{}
	mapstructure.Decode(rawConfig, config)

	// validate every field from the redis config
	if config.Host == "" {
		return nil, errors.New("nil host")
	}
	if config.Port <= 0 {
		return nil, errors.New("nil port")
	}
	if config.Password == "" {
		return nil, errors.New("nil password")
	}
	if config.Database <= 0 {
		return nil, errors.New("nil database")
	}

	// creates a new redis client
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.Database,
	})

	if _, err := client.Ping().Result(); err != nil {
		return nil, errors.New("Invalid redis connection")
	}

	return client, nil
}
