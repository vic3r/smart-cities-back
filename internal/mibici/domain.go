package mibici

import (
	"github.com/gin-gonic/gin"
)

// Domain is a mibici domain
type Domain struct {
	controller *Controller
}

// NewDomain creates a new mibici domain
func NewDomain(usecases UseCases) (*Domain, error) {
	controller, err := newController(usecases)
	if err != nil {
		return nil, err
	}

	return &Domain{controller}, nil
}

// Controller specifies each router pathprefix and each endpoint function
func (d *Domain) Controller(router *gin.Engine) {
	router.GET("mibici/neighborhoods", d.controller.GetNeighborhoods)
	router.GET("mibici/neighborhood/:id", d.controller.GetNeighborhood)
}
