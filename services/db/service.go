package db

import "github.com/go-redis/redis"

// Service is an interface for the redis client implementation
type Service interface {
	Connection(config map[string]interface{}) (*redis.Client, error)
}
