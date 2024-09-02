package main

import (
	"log"
	"net/http"

	"github.com/BarunKGP/nlquery/controllers"
	"github.com/BarunKGP/nlquery/utils"
	"github.com/julienschmidt/httprouter"
)

func enableCors(w *http.ResponseWriter) {
	// Replace this with env var?
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

}

type Signin struct {
	User    any    `json:user`
	Account string `json:account`
}

type User struct {
	Id       string `json: id`
	Name     string `json: name`
	Email    string `json: email`
	ImageSrc string `json: image`
}

func main() {
	logger := utils.CreateLogger()
	router := httprouter.New()

	router.GET("/", controllers.HandleHome(logger))
	router.POST("/api/v1/auth/signin", controllers.HandleSignin(logger))

	logger.Info("Starting http server")
	log.Fatal(http.ListenAndServe(":8001", router))

}
