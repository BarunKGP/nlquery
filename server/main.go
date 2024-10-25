package main

import (
	"log"

	"github.com/BarunKGP/nlquery/adapters"
	// "github.com/BarunKGP/nlquery/controllers"
	// "github.com/BarunKGP/nlquery/internal"
)

func main() {
	gateway := adapters.NewGateway("/api/v1", []string{"http://localhost:3000"})
	if err := gateway.Init(); err != nil {
		log.Fatal("Could not start gateway")
	}

	// env := internal.InitEnv()
	// router := internal.NewApiRouter("/api/v1")
	//
	// gateway := adapters.NewGateway("api/v1").
	// 	WithOrigins([]string{"http://localhost:3000"}).
	// 	WithEnv(env).
	// 	WithRouter(router)

	// var routes []*adapters.GatewayServeOpts
	// //* Open routes
	// or, err1 := adapters.CreateGateway(
	// 	map[string]internal.ControllerFunc{
	// 		"GET /":     controllers.HandleHome,
	// 		"GET /test": controllers.HandleTest,
	//
	// 		"POST /auth/signin": controllers.HandleSignin,
	// 	},
	// 	false,
	// )
	// if err1 != nil {
	// 	log.Fatal("Unable to parse routes")
	// }
	// routes = append(routes, or...)
	// //* Protected routes
	// pr, err2 := adapters.CreateGateway(
	// 	map[string]internal.ControllerFunc{
	// 		"POST /auth/logout": controllers.HandleLogout,
	//
	// 		"POST /user":    controllers.HandleCreateUser,
	// 		"GET /user/:id": controllers.HandleGetUser,
	// 	},
	// 	true,
	// )
	// if err2 != nil {
	// 	log.Fatal("Unable to parse protected routes")
	// }
	// routes = append(routes, pr...)
	// gateway.Handle(routes)

}
