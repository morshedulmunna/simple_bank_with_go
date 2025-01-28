package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morshedulmunna/simple_bank/api/controllers"
)

// RouteConfig holds all controllers needed for route setup
type RouteConfig struct {
	UserController *controllers.UserController
	RootController *controllers.RootController
}

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, config RouteConfig) {
	// Setup root routes
	setupResourceRoutes(router.Group("/"), map[string]interface{}{
		"": map[string]gin.HandlerFunc{
			"GET":    config.RootController.Welcome,
			"HEALTH": config.RootController.HealthCheck,
		},
	})

	v1 := router.Group("/api/v1")

	// Setup all route groups
	setupResourceRoutes(v1, map[string]interface{}{
		"users": map[string]gin.HandlerFunc{
			"GET":    config.UserController.GetUsers,
			"POST":   config.UserController.CreateUser,
			"GET_ID": config.UserController.GetUser,
		},
	})
}

// setupResourceRoutes configures routes for a resource group
func setupResourceRoutes(rg *gin.RouterGroup, resources map[string]interface{}) {
	for resource, handlers := range resources {
		group := rg.Group("/" + resource)
		if handlerMap, ok := handlers.(map[string]gin.HandlerFunc); ok {
			if handler, exists := handlerMap["GET"]; exists {
				group.GET("", handler)
			}
			if handler, exists := handlerMap["POST"]; exists {
				group.POST("", handler)
			}
			if handler, exists := handlerMap["GET_ID"]; exists {
				group.GET("/:id", handler)
			}
			if handler, exists := handlerMap["HEALTH"]; exists {
				group.GET("/health", handler)
			}
		}
	}
}
