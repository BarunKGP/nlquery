package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/BarunKGP/nlquery/internal"
	"github.com/BarunKGP/nlquery/users"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/julienschmidt/httprouter"
)

func HandleHome(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	fmt.Fprintf(w, "Hello from nlQuery\n")
	return nil
}

func GetAuthProviders(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	authProviders := e.AuthConfig.GetProviders()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string][]string{"authProviders": authProviders})
	if err != nil {
		return internal.NewHttpError("Unable to get auth providers", http.StatusInternalServerError, r.URL.Path)
	}

	return nil
}

func HandleSignin(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
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

func HandleGetUser(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	queries := users.New(e.DB)
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to convert id: %v to int: %v", p.ByName("id"), err.Error())
		httpErr := internal.NewHttpError(errMsg, http.StatusInternalServerError, r.URL.Path)
		e.Logger.Error(httpErr.Error())
		return httpErr
	}

	user, err := queries.GetUser(e.DbCtx, id)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to fetch user %v from db: %v", id, err.Error())
		httpErr := internal.NewHttpError(errMsg, http.StatusInternalServerError, r.URL.Path)
		e.Logger.Error(httpErr.Error())
		return httpErr
	}

	apiUser := internal.ApiUser{
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

	e.Logger.Info(fmt.Sprintf("Received user with id: %d: %+v", id, user))
	return nil

}

func HandleCreateUser(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	queries := users.New(e.DB)

	userParams := users.CreateUserParams{}
	if err := json.NewDecoder(r.Body).Decode(&userParams); err != nil {
		errMsg := fmt.Sprintf("Error decoding body: %v", err.Error())
		httpErr := internal.NewHttpError(errMsg, http.StatusInternalServerError, r.URL.Path)
		e.Logger.Error(httpErr.Error())
		return httpErr
	}
	userParams.Createdat = pgtype.Timestamp{Time: time.Now()}
	userParams.Lastmodified = pgtype.Timestamp{Time: time.Now()}

	user, err := queries.CreateUser(e.DbCtx, userParams)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to create user: %v", err.Error())
		httpErr := internal.HttpStatusError{
			Message: errMsg,
			Status:  http.StatusInternalServerError,
			Path:    r.URL.Path,
		}
		e.Logger.Error(httpErr.Error())
		return httpErr
	}

	e.Logger.Info(fmt.Sprintf("User created successfully: %+v", user))

	userResponse := internal.ApiUser{
		Name:  user.Name,
		Email: user.Email,
		// SessionId: "",
		UserId: string(user.ID),
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userResponse); err != nil {
		return fmt.Errorf("Unable to convert payload: %v to json", userResponse)

	}

	return nil
}
