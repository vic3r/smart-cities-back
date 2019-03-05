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

	return
}

// GetNeighborhood returns a json with a particular neighborhood
func (c *Controller) GetNeighborhood(ctx *gin.Context) {
	zoneName := ctx.Param("zone_name")
	neighborhoodID := ctx.Param("neighborhood_id")

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

//GetNeighborhoodsByZone returns a json with a particular zone
func (c *Controller) GetNeighborhoodsByZone(ctx *gin.Context) {
	zoneName := ctx.Param("zone_name")

	if zoneName == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "nil zone_id"})
		ctx.Abort()

		return
	}
	// TODO: handle errors properly
	neighborhoods, err := c.usecases.GetNeighborhoodsByZone(zoneName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"no neighborhoods found": err})

		return
	}

	messageResponse := fmt.Sprintf("neighborhoods in %s", zoneName)
	ctx.JSON(http.StatusOK, gin.H{messageResponse: neighborhoods})

	return
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

	return
}
