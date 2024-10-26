package controllers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/BarunKGP/nlquery/adapters"
	"github.com/BarunKGP/nlquery/core/auth"
	"github.com/BarunKGP/nlquery/ports"
	"github.com/julienschmidt/httprouter"
)

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
}
