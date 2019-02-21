package mibici

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type controller struct {
	usecases UseCases
}

// NewController creates a controller
func newController(usecases UseCases) (*controller, error) {
	if usecases == nil {
		return nil, errors.New("nil usecases")
	}

	return &controller{usecases}, nil
}

func (c *controller) GetNeighborhoods(ctx *gin.Context) {
	neighborhoods, err := c.usecases.GetNeighborhoods()
	if err != nil {
		// TODO: handle errors properly
		// implement a switch for different errors
		ctx.JSON(404, gin.H{"no neighborhoods found": err})
	}

	ctx.JSON(200, gin.H{"neighborhoods": neighborhoods})

	return
}

func (c *controller) GetNeighborhood(ctx *gin.Context) {
	neighborhoodID := ctx.Param("neighborhood_id")

	if neighborhoodID == "" {
		ctx.JSON(500, gin.H{"error": "nil neighborhood_id"})
		ctx.Abort()
		return
	}
	// TODO: handle errors properly
	neighborhood, err := c.usecases.GetNeighborhood(neighborhoodID)
	if err != nil {
		ctx.JSON(404, gin.H{"no neighborhood found": err})
	}

	ctx.JSON(200, gin.H{"neighborhood": neighborhood})
	return
}
