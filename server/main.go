package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func enableCors(w *http.ResponseWriter) {
	// Replace this with env var?
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

}

type Signin struct {
	User any `json:user`
	Account string `json:account`
}

type User struct {
	Id string `json: id`
	Name string `json: name`
	Email string `json: email`
	ImageSrc string `json: image`
}

func main() {
	logger := CreateLogger()
	router := httprouter.New()
	
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		logger.Info("Home route")
		fmt.Fprintf(w, "Hello from nlQuery!\n")
	})

	router.POST("/api/v1/auth/signin", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		body, err := io.ReadAll(r.Body); if err != nil {
			errMsg := "Unable to read body"
			logger.Error(errMsg, err)
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("X-Error", errMsg)

			return
		}

		var signinBody Signin
		err = json.Unmarshal(body, &signinBody); if err != nil {
				errMsg := "Unable to decode signin body"
				logger.Error(errMsg, err)
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Add("X-Error", errMsg)

				return
		}


		// logger.Info("On signin route")
		// logger.Debug(str(signinBody))
	})



	logger.Info("Starting http server")
	log.Fatal(http.ListenAndServe(":8001", router))

}

