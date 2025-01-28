package controllers

import (
	"github.com/gin-gonic/gin"
)

// RootController handles root-level routes
type RootController struct{}

// NewRootController creates a new RootController instance
func NewRootController() *RootController {
	return &RootController{}
}

// Welcome handles the root welcome route
func (rc *RootController) Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Simple Bank API",
	})
}

// HealthCheck handles the health check route
func (rc *RootController) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
	})
}
