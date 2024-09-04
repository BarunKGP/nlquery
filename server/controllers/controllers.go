package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HandleHome(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	fmt.Fprintf(w, "Hello from nlQuery\n")
	return nil
}

func HandleSignin(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		errMsg := "Unable to read body"
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("X-Error", errMsg)

		return fmt.Errorf(`Error handling signin: %s`, errMsg)
	}

	fmt.Printf("response body: %v\n", body)
	return nil

	// Create session

	// Write signed in user details to db

	// Send session details to frontend
}
