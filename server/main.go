package main

import (
	"log"

	"github.com/BarunKGP/nlquery/adapters"
	"github.com/BarunKGP/nlquery/adapters/controllers"
)

func main() {
	gateway := adapters.NewGateway("/api/v1", []string{"http://localhost:3000"})
	routes := controllers.InitRoutes()
	if err := gateway.Init(routes); err != nil {
		log.Fatalf("Fatal! Initialization failed: %v", err)
	}
}
