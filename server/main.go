package main

import (
	"log"

	"github.com/BarunKGP/nlquery/adapters"
	"github.com/BarunKGP/nlquery/adapters/controllers"
	// "github.com/BarunKGP/nlquery/ports"
)

func main() {
	gateway := adapters.NewGateway("/api/v1", []string{"http://localhost:3000"})
	routes := controllers.InitRoutes()
	// var _ ports.Routeable = routes
	if err := gateway.Init(routes); err != nil {
		log.Fatalf("Fatal! Initialization failed: %v", err)
	}
}
