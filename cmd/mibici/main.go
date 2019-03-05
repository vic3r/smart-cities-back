package main

import (
	"flag"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/vic3r/smart-cities-back/services/db/redis"
)

var configurations = map[string]string{
	"DEBUG":      "debug",
	"DEV":        "development",
	"PRODUCTION": "production",
	"LOCAL":      "local",
}

var (
	configPath,
	configType string
)

func init() {
	// configPath describes the configuration path to find
	flag.StringVar(&configPath, "ConfigDir", "", "directory for the configuration files")
	// ConfigType describes the configuration type to read
	flag.StringVar(&configType, "ConfigType", "", "configuration type to look for")
}

func main() {
	// Get the config environment
	configEnv := os.Getenv("ENVIRONMENT")

	// check if the current config env exists
	configName, ok := configurations[configEnv]
	if !ok {
		log.Fatal("invalid config environment")
	}

	flag.Parse()

	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
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
