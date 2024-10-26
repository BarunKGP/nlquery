package controllers

import (
	// "context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/BarunKGP/nlquery/adapters"
	"github.com/BarunKGP/nlquery/ports"
	"github.com/julienschmidt/httprouter"
)

func HandleHome(e ports.EnvReader, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Hello from nlQuery!"}); err != nil {
		return fmt.Errorf("Error encoding message")
	}
	slog.Info("Hello from nlQuery\n")
	return nil
}

func HandleTest(e ports.EnvReader, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	w.Header().Set("Content-Type", "application/json")
	path := r.URL.Path
	host := r.URL.Hostname()
	if err := json.
		NewEncoder(w).
		Encode(map[string]string{"message": "Test route", "path": path, "host": host}); err != nil {
		return fmt.Errorf("Error encoding message")
	}
	slog.Debug("Hello from nlQuery: Test route\n")
	return nil
}

func HandleGetUser(e ports.EnvReader, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	queries := e.GetDbFactory()
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to convert id: %v to int: %v", p.ByName("id"), err.Error())
		httpErr := adapters.NewHttpError(errMsg, http.StatusInternalServerError, r.URL.Path)
		slog.Error(httpErr.Error())
		return httpErr
	}

	user, err := queries.GetUser(r.Context(), id)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to fetch user %v from db: %v", id, err.Error())
		httpErr := adapters.NewHttpError(errMsg, http.StatusInternalServerError, r.URL.Path)
		slog.Error(httpErr.Error())
		return httpErr
	}

	apiUser := adapters.ApiUser{
		Name:  user.Name,
		Email: user.Email,
		// SessionId: "", // TODO: figure out auth
		UserId: string(user.ID),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(apiUser); err != nil {
		return fmt.Errorf("Unable to convert payload: %v to json", apiUser)
	}

	slog.Info(fmt.Sprintf("Received user with id: %d: %+v", id, user))
	return nil
}

func HandleCreateUser(e ports.EnvReader, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	slog.Debug("Reached HandleCreateUser")
	dbFactory := e.GetDbFactory()
	apiUser, err := adapters.NewApiUser().FromHttp(r)
	if err != nil {
		errMsg := fmt.Sprintf("Error decoding body: %v", err.Error())
		httpErr := adapters.NewHttpError(errMsg, http.StatusInternalServerError, r.URL.Path)
		slog.Error(httpErr.Error())
		return httpErr
	}
	slog.Info("Received api user: ", slog.Any("apiUser", apiUser))

	if _, err = apiUser.Persist(r.Context(), dbFactory); err != nil {
		slog.Error("Could not create user", "user(API)", apiUser, "error", err)
		return fmt.Errorf("Could not persist user")
	}
	return err
}

func writeObjectToJson(w http.ResponseWriter, v any) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("Unable to convert payload: %v to json", v)
	}

	return nil
}
