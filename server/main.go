package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func enableCors(w *http.ResponseWriter) {
	// Replace this with env var?
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

}

func main() {
	logger := CreateLogger()
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		logger.Info("Home route")
		fmt.Fprintf(w, "Hello from nlQuery!\n")
	})

	logger.Info("Starting http server")
	log.Fatal(http.ListenAndServe(":8001", router))

}
