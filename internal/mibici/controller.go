package mibici

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is a public controller in where is handled every gin endpoint
type Controller struct {
	usecases UseCases
}

// NewController creates a Controller
func newController(usecases UseCases) (*Controller, error) {
	if usecases == nil {
		return nil, errors.New("nil usecases")
	}

	return &Controller{usecases}, nil
}

// GetStation returns a json with a particular station
func (c *Controller) GetStation(ctx *gin.Context) {
	zoneName := ctx.Param("zone_name")
	stationID := ctx.Param("id")
	if stationID == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "nil station_id"})
		ctx.Abort()
		return
	}

	if zoneName == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "nil zone_name"})
		ctx.Abort()
		return
	}

	station, err := c.usecases.GetStation(zoneName, stationID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"no station found": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"station": station})
}

// GetNeighborhoods returns a json with a list of neighborhoods
func (c *Controller) GetNeighborhoods(ctx *gin.Context) {
	neighborhoods, err := c.usecases.GetNeighborhoods()
	if err != nil {
		// TODO: handle errors properly
		// implement a switch for different errors
		ctx.JSON(http.StatusNotFound, gin.H{"no neighborhoods found": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"neighborhoods": neighborhoods})
}

// GetNeighborhood returns a json with a particular neighborhood
func (c *Controller) GetNeighborhood(ctx *gin.Context) {
	zoneName := ctx.Param("zone_name")
	neighborhoodID := ctx.Param("name")

	if neighborhoodID == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "nil neighborhood_id"})
		ctx.Abort()
		return
	}

	if zoneName == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "nil zone_name"})
		ctx.Abort()
		return
	}

	// TODO: handle errors properly
	neighborhood, err := c.usecases.GetNeighborhood(zoneName, neighborhoodID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"no neighborhood found": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"neighborhood": neighborhood})
}

//GetStationsByZone returns a json with a particular zone
func (c *Controller) GetStationsByZone(ctx *gin.Context) {
	zoneName := ctx.Param("zone_name")

	if zoneName == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "nil zone_id"})
		ctx.Abort()
		return
	}
	// TODO: handle errors properly
	stations, err := c.usecases.GetStationsByZone(zoneName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"no stations found": err})
		return
	}

	messageResponse := fmt.Sprintf("stations in %s", zoneName)
	ctx.JSON(http.StatusOK, gin.H{messageResponse: stations})
}

//GetZones returns a json with all zones
func (c *Controller) GetZones(ctx *gin.Context) {

	// TODO: handle errors properly
	zones, err := c.usecases.GetZones()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"no zones found": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"zones": zones})
}
