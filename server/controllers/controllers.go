package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/BarunKGP/nlquery/internal"
	"github.com/BarunKGP/nlquery/internal/auth"
	"github.com/BarunKGP/nlquery/users"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/julienschmidt/httprouter"
	"github.com/markbates/goth/gothic"
)

func HandleHome(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	fmt.Fprintf(w, "Hello from nlQuery\n")
	return nil
}

func GetAuthProviders(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	authProviders := auth.GetProviders()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string][]string{"authProviders": authProviders})
	if err != nil {
		return internal.NewHttpError("Unable to get auth providers", http.StatusInternalServerError, r.URL.Path)
	}

	return nil
}

func HandleSignin(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	provider := p.ByName("provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		gothic.BeginAuthHandler(w, r)
		return nil
	}

	user := internal.ApiUser{
		Name:   gothUser.Name,
		Email:  gothUser.Email,
		UserId: gothUser.UserID,
	}

	e.Logger.Info(fmt.Sprintf("Signin by user: %+v", user))

	// Check if user exists in db
	// TODO: How should we search for user?
	//! UserId returned by provider will be different from our own
	queries := users.New(e.DB)
	queryId, convErr := strconv.ParseInt(gothUser.UserID, 10, 64)
	if convErr != nil {
		return fmt.Errorf("Error converting goth user id: %v to int64", gothUser.UserID)
	}

	rawUser, _err := queries.GetUser(e.DbCtx, queryId)
	if _err != nil {
		// User not found. Create new user
		userParams := users.CreateUserParams{
			Name:         user.Name,
			Email:        user.Email,
			Provider:     pgtype.Text{String: "google"},
			Createdat:    pgtype.Timestamp{Time: time.Now()},
			Lastmodified: pgtype.Timestamp{Time: time.Now()},
		}

		user, err := queries.CreateUser(e.DbCtx, userParams)
		if err != nil {
			return internal.HttpStatusError{
				Message: fmt.Sprintf("Unable to create user: %v", err.Error()),
				Status:  http.StatusInternalServerError,
				Path:    r.URL.Path,
			}
		}

		e.Logger.Info(fmt.Sprintf("User created successfully: %+v", user))
	}

	e.Logger.Info(fmt.Sprintf("User exists: %+v", rawUser))

	// TODO: Create session

	// TODO: Write signed in user details to db

	// TODO: Send session details to frontend

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		return fmt.Errorf("Error converting user: %v to API JSON response", gothUser)
	}

	http.Redirect(w, r, fmt.Sprint("%s/user/%s", e.ClientAuthRedirect, gothUser.UserID), http.StatusFound)

	return nil
}

func HandleAuthCallback(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	provider := p.ByName("provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return internal.NewHttpError("Unable to complete authentication", http.StatusInternalServerError, r.URL.Path)

	}

	user := internal.ApiUser{
		Name:   gothUser.Name,
		Email:  gothUser.Email,
		UserId: gothUser.UserID,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		return fmt.Errorf("Error converting user: %v to API JSON response", gothUser)
	}

	http.Redirect(w, r, fmt.Sprint("%s/user/%s", e.ClientAuthRedirect, gothUser.UserID), http.StatusFound)

	return nil

}

func HandleLogout(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	provider := p.ByName("provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.Logout(w, r)
	http.Redirect(w, r, fmt.Sprint("%s/", e.ClientAuthRedirect), http.StatusTemporaryRedirect)

	return nil
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
	e.Logger.Info("Reached HandleCreateUser")
	queries := users.New(e.DB)

	apiUser := internal.ApiUser{}
	if err := json.NewDecoder(r.Body).Decode(&apiUser); err != nil {
		errMsg := fmt.Sprintf("Error decoding body: %v", err.Error())
		httpErr := internal.NewHttpError(errMsg, http.StatusInternalServerError, r.URL.Path)
		e.Logger.Error(httpErr.Error())
		return httpErr
	}

	e.Logger.Info("Received api user: ", slog.Any("apiUser", apiUser))

	userParams := users.CreateUserParams{
		Name:         apiUser.Name,
		Email:        apiUser.Email,
		Provider:     pgtype.Text{String: "google", Valid: true},
		Createdat:    pgtype.Timestamp{Time: time.Now().UTC(), InfinityModifier: pgtype.Finite, Valid: true},
		Lastmodified: pgtype.Timestamp{Time: time.Now().UTC(), InfinityModifier: pgtype.Finite, Valid: true},
	}

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
