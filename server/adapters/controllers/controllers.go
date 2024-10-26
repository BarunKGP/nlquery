package controllers

import (
	// "context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/BarunKGP/nlquery/adapters"
	"github.com/BarunKGP/nlquery/core/auth"
	"github.com/BarunKGP/nlquery/ports"
	"github.com/julienschmidt/httprouter"
	// "github.com/markbates/goth/gothic"
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

func HandleSignin(e ports.EnvReader, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	user, err := adapters.NewApiUser().FromHttp(r)
	if err != nil {
		return adapters.NewHttpError("Unable to parse request body", http.StatusBadRequest, r.URL.Path)
	}
	slog.Info("Signin attempt", "user", user)

	dbFactory := e.GetDbFactory()
	if ok := user.IsPersisted(r.Context(), dbFactory); !ok {
		slog.Info("New signin", "user", user)
		userPersisted, err := user.Persist(r.Context(), dbFactory)
		if err != nil {
			slog.Error("Failed to persist user in database", "user", user)
			return fmt.Errorf("Failed to persist user")
		}
		slog.Info("Persisted new signin user", "user", userPersisted)
	}
	slog.Info("Returning signin", "user", user)

	// Create auth token
	token, err := auth.CreateToken(user.UserId)
	if err != nil {
		return fmt.Errorf("Error creating token: %v", err)
	}
	slog.Info(fmt.Sprintf("Returning JWT: %v", token))

	// TODO: Write signed in user details to db

	// Send token back to frontend through a cookie
	cookie := http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   false, // TODO: change to true in prod!,
		Expires:  time.Now().AddDate(1, 0, 0),
	}
	http.SetCookie(w, &cookie)
	slog.Info(fmt.Sprintf("Cookie written successfully: %+v", cookie))

	return nil
}

// func HandleAuthCallback(e ports.EnvReader, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
// 	provider := p.ByName("provider")
// 	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
//
// 	gothUser, err := gothic.CompleteUserAuth(w, r)
// 	if err != nil {
// 		return adapters.NewHttpError("Unable to complete authentication", http.StatusInternalServerError, r.URL.Path)
// 	}
//
// 	user := core.ApiUser{
// 		Name:   gothUser.Name,
// 		Email:  gothUser.Email,
// 		UserId: gothUser.UserID,
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(user); err != nil {
// 		return fmt.Errorf("Error converting user: %v to API JSON response", gothUser)
// 	}
//
// 	http.Redirect(w, r, fmt.Sprint("%s/user/%s", e.ClientAuthRedirect, gothUser.UserID), http.StatusFound)
//
// 	return nil
// }

func HandleLogout(e ports.EnvReader, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
	return nil

	// provider := p.ByName("provider")
	// r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	// gothic.Logout(w, r)
	// http.Redirect(w, r, fmt.Sprint("%s/", e.ClientAuthRedirect), http.StatusTemporaryRedirect)
	//

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
