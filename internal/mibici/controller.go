package mibici

import (
	"errors"
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
	}

	ctx.JSON(http.StatusOK, gin.H{"neighborhoods": neighborhoods})

	return
}

// GetNeighborhood returns a json with a particular neighborhood
func (c *Controller) GetNeighborhood(ctx *gin.Context) {
	neighborhoodID := ctx.Param("neighborhood_id")

	if neighborhoodID == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "nil neighborhood_id"})
		ctx.Abort()

		return
	}
	// TODO: handle errors properly
	neighborhood, err := c.usecases.GetNeighborhood(neighborhoodID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"no neighborhood found": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"neighborhood": neighborhood})

	return
}

//GetNeighborhoodByZone returns a json with a particular zone
func (c *Controller) GetNeighborhoodByZone(ctx *gin.Context) {
	zoneID := ctx.Param("zone_id")

	if zoneID == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "nil zone_id"})
		ctx.Abort()

		return
	}
	// TODO: handle errors properly
	zone, err := c.usecases.GetNeighborhoodByZone(zoneID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"no zone found": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"zone": zone})

	return
}
