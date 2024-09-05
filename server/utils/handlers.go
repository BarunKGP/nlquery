package utils

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/julienschmidt/httprouter"
)

type IApiError interface {
	error
	GetStatus() int
}

type HttpStatusError struct {
	Message error
	status  int
	path    string
}

func (err HttpStatusError) Error() string {
	return fmt.Sprintf("Error: %v at %v: returning HTTP %d",
		err.Message, err.path, err.status)
}

func (err HttpStatusError) GetStatus() int {
	return err.status
}

func (err HttpStatusError) GetPath() string {
	return err.path
}

type Env struct {
	DB     *pgx.Conn
	Port   string
	Host   string
	Logger *slog.Logger
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
}
