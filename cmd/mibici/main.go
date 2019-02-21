package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/vic3r/smart-cities-back/services/db/redis"
)

const (
	// ConfigPath describes the configuration file path to look for configuration files
	ConfigPath = "$HOME/go/src/github.com/vic3r/smart-cities-back/config"
	// ConfigType describes the configuration type to read
	ConfigType = "yaml"
)

var configurations = map[string]string{
	"DEBUG":      "debug",
	"DEV":        "development",
	"PRODUCTION": "production",
}

func main() {
	// Get the config environment
	configEnv := os.Getenv("ENVIRONMENT")
	viper.AddConfigPath(ConfigPath)
	viper.SetConfigType(ConfigType)

	// check if the current config env exists
	configName, ok := configurations[configEnv]
	if !ok {
		log.Fatal("invalid config environment")
	}

	viper.SetConfigFile(configName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config file: %s \n", err)
	}

	// create redis db service
	dbService, err := redis.New()
	if err != nil {
		log.Fatalf("error creating db service: %s \n", err)
	}

	// create a mibici domain
	mibiciDomain, err := createMibiciDomain(dbService)
	if err != nil {
		log.Fatal(err)
	}

	runServer(mibiciDomain)
}
