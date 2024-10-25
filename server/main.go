package main

import (
	"log"

	"github.com/BarunKGP/nlquery/adapters"
)

func main() {
	gateway := adapters.NewGateway("/api/v1", []string{"http://localhost:3000"})
	if err := gateway.Init(); err != nil {
		log.Fatalf("Could not start gateway: %v", err)
	}
}
