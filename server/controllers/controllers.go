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
	"github.com/BarunKGP/nlquery/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/julienschmidt/httprouter"
	"github.com/markbates/goth/gothic"
)

func HandleHome(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Hello from nlQuery!"}); err != nil {
		return fmt.Errorf("Error encoding message")
	}

	e.Logger.Info("Hello from nlQuery\n")

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
	var user internal.ApiUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return internal.NewHttpError("Unable to parse request body", http.StatusBadRequest, r.URL.Path)
	}

	e.Logger.Info(fmt.Sprintf("Signin by user: %+v", user))

	// Check if user exists in db
	queries := database.New(e.DB)

	// if _, err := queries.GetUserByProviderUserId(e.DbCtx, pgtype.Text{String: user.UserId, Valid: true}); err != nil {
	if _, err := queries.GetUserByEmail(e.DbCtx, user.Email); err != nil {
		// User not found. Create new user
		userParams := database.CreateUserParams{
			Name:           user.Name,
			Email:          user.Email,
			Imagesrc:       pgtype.Text{String: user.ImageSrc, Valid: true},
			Provideruserid: pgtype.Text{String: user.UserId, Valid: true},
			Createdat:      pgtype.Timestamptz{Time: time.Now().UTC(), InfinityModifier: pgtype.Finite, Valid: true},
			Lastmodified:   pgtype.Timestamptz{Time: time.Now().UTC(), InfinityModifier: pgtype.Finite, Valid: true},
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

	// Create token
	token, err := auth.CreateToken(user.UserId)
	if err != nil {
		return fmt.Errorf("Error creating token: %v", err)
	}

	e.Logger.Info(fmt.Sprintf("Returning JWT: %v", token))

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
	e.Logger.Info("Cookie written successfully")

	// Send token back to frontend through response
	// w.Header().Set("Content-Type", "application/json")
	// if err := json.NewEncoder(w).Encode(map[string]string{"token": token}); err != nil {
	// 	return fmt.Errorf("Error converting token: %v to API JSON response", token)
	// }

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
	// provider := p.ByName("provider")
	// r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})

	// gothic.Logout(w, r)
	// http.Redirect(w, r, fmt.Sprint("%s/", e.ClientAuthRedirect), http.StatusTemporaryRedirect)
	//
	return nil
}

func HandleGetUser(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	queries := database.New(e.DB)
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
	queries := database.New(e.DB)

	apiUser := internal.ApiUser{}
	if err := json.NewDecoder(r.Body).Decode(&apiUser); err != nil {
		errMsg := fmt.Sprintf("Error decoding body: %v", err.Error())
		httpErr := internal.NewHttpError(errMsg, http.StatusInternalServerError, r.URL.Path)
		e.Logger.Error(httpErr.Error())
		return httpErr
	}

	e.Logger.Info("Received api user: ", slog.Any("apiUser", apiUser))

	// Search for existing user
	queryId, convErr := strconv.Atoi(apiUser.UserId)
	if convErr != nil {
		return fmt.Errorf("Error converting user id: %v to int64", apiUser.UserId)
	}

	if _, err := queries.GetUser(e.DbCtx, int64(queryId)); err != nil {
		// User not found - create new user
		userParams := database.CreateUserParams{
			Name:           apiUser.Name,
			Email:          apiUser.Email,
			Imagesrc:       pgtype.Text{String: apiUser.ImageSrc, Valid: true},
			Provideruserid: pgtype.Text{String: apiUser.UserId, Valid: true},
			Createdat:      pgtype.Timestamptz{Time: time.Now().UTC(), InfinityModifier: pgtype.Finite, Valid: true},
			Lastmodified:   pgtype.Timestamptz{Time: time.Now().UTC(), InfinityModifier: pgtype.Finite, Valid: true},
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
		} else {
			e.Logger.Info(fmt.Sprintf("User with provider ID: %s exists", user.Provideruserid))
		}

		e.Logger.Info(fmt.Sprintf("User created successfully: %+v", user))
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "success"}); err != nil {
		return fmt.Errorf("Unable to return success message")
	}

	return nil
}

func writeObjectToJson(w http.ResponseWriter, v any) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("Unable to convert payload: %v to json", v)
	}

	return nil
}
