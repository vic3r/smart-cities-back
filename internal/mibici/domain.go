package mibici

import (
	"github.com/gin-gonic/gin"
)

// Domain is a mibici domain
type Domain struct {
	controller *Controller
}

// NewDomain creates a new mibici domain
func NewDomain(usecases UseCases, populatorURL string) (*Domain, error) {
	controller, err := newController(usecases, populatorURL)
	if err != nil {
		return nil, err
	}

	return &Domain{controller}, nil
}

// Controller specifies each router pathprefix and each endpoint function
func (d *Domain) Controller(router *gin.Engine) {
	router.GET("mibici/zone/:zone_name/station/:id", d.controller.GetStation)
	router.GET("mibici/neighborhoods", d.controller.GetNeighborhoods)
	router.GET("mibici/zone/:zone_name/neighborhood/:id", d.controller.GetNeighborhood)
	router.GET("mibici/stations/zone/:zone_name", d.controller.GetStationsByZone)
	router.GET("mibici/zones", d.controller.GetZones)
}
