package fake

import (
	"github.com/go-redis/redis"
	"github.com/vic3r/smart-cities-back/services/db"
)

type fake struct{}

var _ db.Service = &fake{}

// Connection returns a fake connection
func (f *fake) Connection(rawConfig map[string]interface{}) (*redis.Client, error) {
	return nil, nil
}
