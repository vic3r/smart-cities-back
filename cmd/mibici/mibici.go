package main

import (
	"log"

	"github.com/vic3r/smart-cities-back/services/db"

	"github.com/spf13/viper"
	mibici "github.com/vic3r/smart-cities-back/internal/mibici"
	"github.com/vic3r/smart-cities-back/internal/mibici/usecases"
	storage "github.com/vic3r/smart-cities-back/internal/mibici/usecases/storage/factory"
)

func createMibiciDomain(dbService db.Service) (*mibici.Domain, error) {

	// read mibici storage config
	storageConfig := viper.Get("domains.mibici.storage").(map[string]interface{})
	if storageConfig == nil {
		log.Fatal("error loading mibici db config \n")
	}

	// create mibici storage instance
	mibiciStorage, err := storage.New(storageConfig, dbService)
	if err != nil {
		log.Fatalf("error creating mibici storage instance: %v \n", err)
	}

	// create mibici use cases instance
	mibiciUsecases, err := usecases.New(mibiciStorage)
	if err != nil {
		log.Fatalf("error creating mibici usecases instance: %v \n", err)
	}

	// read mibici storage config
	populatorURL := viper.Get("domains.mibici.populator.url").(string)
	if populatorURL == "" {
		log.Fatal("error loading populator url \n")
	}

	// create mibici domain instance
	mibiciDomain, err := mibici.NewDomain(mibiciUsecases, populatorURL)
	if err != nil {
		log.Fatalf("error creating mibici domain instance: %v \n", err)
	}

	return mibiciDomain, nil
}
