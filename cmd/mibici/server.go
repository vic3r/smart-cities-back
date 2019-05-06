package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"github.com/vic3r/smart-cities-back/internal/mibici"
)

func runServer(mibiciDomain *mibici.Domain) {
	// create a new gin controller and handle it
	router := gin.Default()
	router.Use(gin.Recovery())
	mibiciDomain.Controller(router)

	// read host and port from config
	host := viper.Get("api.host").(string)
	port := viper.Get("api.port").(int)

	// create address and run the server
	address := fmt.Sprintf("%s:%d", host, port)
	handler := cors.Default().Handler(router)

	http.ListenAndServe(address, handler)
}
