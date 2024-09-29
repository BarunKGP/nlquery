package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	// "github.com/BarunKGP/nlquery/internal/auth"
	"github.com/BarunKGP/nlquery/internal/auth"
	"github.com/jackc/pgx/v5"
	"github.com/julienschmidt/httprouter"
)

type IApiError interface {
	Error() string
	GetStatus() int
}

type HttpStatusError struct {
	Message string
	Status  int
	Path    string
}

func (err HttpStatusError) Error() string {
	return fmt.Sprintf("Error: %v at %v: returning HTTP %d",
		err.Message, err.Path, err.Status)
}

func (err HttpStatusError) GetStatus() int {
	return err.Status
}

func (err HttpStatusError) GetPath() string {
	return err.Path
}

func NewHttpError(errMsg string, status int, path string) HttpStatusError {
	// Build a new HttpStatusError
	// This can be written to a logger using `httpErr.Error()` or returned as an error value
	return HttpStatusError{
		Message: errMsg,
		Status:  status,
		Path:    path,
	}
}

type Env struct {
	DB                 *pgx.Conn
	Port               string
	Host               string
	Logger             *slog.Logger
	DbCtx              context.Context
	ClientAuthRedirect string
}

type ControllerFunc func(e *Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error

func (e *Env) Handle(fn ControllerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		logger := e.Logger
		// CORS
		w.Header().Add("Access-Control-Allow-Origin", e.ClientAuthRedirect)
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		err := fn(e, w, r, p)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Content-Type-Options", "nosniff")

			switch err := err.(type) {
			case IApiError:
				logger.Error(err.Error())

				//* This could be its own function
				w.WriteHeader(err.GetStatus())
				if err := json.NewEncoder(w).Encode(map[string]string{"message": err.Error()}); err != nil {
					log.Fatal("Error encoding JSON response")
				}

			default:
				logger.Error(fmt.Sprintf("Internal error occurred: %v", err.Error()))

				w.WriteHeader(http.StatusInternalServerError)
				if err := json.NewEncoder(w).Encode(map[string]string{"message": "Uh oh... we need a minute :("}); err != nil {
					log.Fatal("Error encoding JSON response")
				}
			}
		}
	}
}

func (e *Env) HandleProtected(fn ControllerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		logger := e.Logger
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			logger.Error(fmt.Sprintf("Unauthorized: %v", err))
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		tokenString := cookie.Value
		if err := auth.VerifyToken(tokenString); err != nil {
			logger.Error(fmt.Sprintf("Unauthorized: %v", err))
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		err = fn(e, w, r, p)
		if err != nil {
			switch err := err.(type) {
			case IApiError:
				logger.Error(err.Error())
				http.Error(w, err.Error(), err.GetStatus())
			default:
				logger.Error(fmt.Sprintf("Internal error occurred: %v", err.Error()))
				http.Error(w, "Uh oh... we need a minute :(", http.StatusInternalServerError)
			}
		}
	}
}

type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := h.H(h.Env, w, r, p)
	if err != nil {
		logger := h.Logger
		switch e := err.(type) {
		case IApiError:
			logger.Error(e.Error())
			http.Error(w, e.Error(), e.GetStatus())
		default:
			logger.Error(fmt.Sprintf("Unknown error: %v", e.Error()))
			http.Error(w, "Unknown error occurred", http.StatusInternalServerError)
		}
	}
}

func Handle(h Handler) httprouter.Handle {
	return h.ServeHTTP
	// err := h.H(h.Env, w, r, p)
	// if err != nil {
	// 	logger := h.Logger
	// 	switch e := err.(type) {
	// 	case IApiError:
	// 		logger.Error(e.Error())
	// 		http.Error(w, e.Error(), e.GetStatus())
	// 	default:
	// 		logger.Error(fmt.Sprintf("Unknown error: %v", e.Error()))
	// 		http.Error(w, "Unknown error occurred", http.StatusInternalServerError)
	// 	}
	// }
}
