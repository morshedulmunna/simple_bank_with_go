package main

import (
	"github.com/morshedulmunna/simple_bank/api"
	"github.com/morshedulmunna/simple_bank/api/controllers"
	"github.com/morshedulmunna/simple_bank/api/routes"
)

func main() {
	config := routes.RouteConfig{
		UserController: &controllers.UserController{},
	}
	server := api.NewServer(config)

	// Start the server on port 8080
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
