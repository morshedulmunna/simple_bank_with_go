package api

import (
	"github.com/gin-gonic/gin"
	"github.com/morshedulmunna/simple_bank/api/middleware"
	"github.com/morshedulmunna/simple_bank/api/routes"
)

type Server struct {
	router      *gin.Engine
	routeConfig routes.RouteConfig
}

func NewServer(config routes.RouteConfig) *Server {
	server := &Server{
		routeConfig: config,
	}
	router := gin.Default()

	// Configure CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Setup routes
	routes.SetupRoutes(router, config)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
