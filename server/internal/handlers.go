package internal

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

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
	// 	This can be written to a logger using `httpErr.Error()` or returned as an error value
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
		err := fn(e, w, r, p)
		if err != nil {
			logger := e.Logger
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
