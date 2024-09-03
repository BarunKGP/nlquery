package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func enableCors(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		next(w, r, p)
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello from nlQuery\n")
}

func HandleSignin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errMsg := "Unable to read body"
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("X-Error", errMsg)
	}

	// Create session

	// Write signed in user details to db

	// Send session details to frontend
}
