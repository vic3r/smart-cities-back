package factory

import (
	"fmt"

	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage"
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage/fake"
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage/redis"
	"github.com/vic3r/smart-cities-back/services/db"
)

const (
	redisID = "redis"
)

// New creates a new mibici storage based on an implementation and config
func New(rawConfig map[string]interface{}, dbService db.Service) (storage.Storage, error) {
	implementation := rawConfig["implementation"].(string)
	if implementation == redisID {
		db, err := dbService.Connection(rawConfig)
		if err != nil {
			return nil, fmt.Errorf("error creating %s storage: %v", redisID, err)
		}

		return redis.New(db)
	}

	return fake.New()
}
