package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/BarunKGP/nlquery/internal/auth"
	"github.com/BarunKGP/nlquery/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

type Env struct {
	DB     database.DBTX
	Port   uint16
	Host   string
	Logger *slog.Logger
	DbCtx  context.Context
}

func getDbUrl() string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}

func InitEnv() *Env {
	if e := godotenv.Load(); e != nil {
		log.Fatal("Unable to read environment variables")
	}
	dbUrl := getDbUrl()
	// TODO: Change level based on env: debug for dev, info for stg, warn for prod
	logger := CreateLogger(slog.LevelDebug)
	logger.Debug("Connecting to postgres db with dbUrl: " + dbUrl)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer conn.Close(ctx)

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = "localhost"
	}
	// TODO: Make this programmatic
	auth.NewAuthConfig([]string{"google", "github"})

	var port uint16
	if num64, err := strconv.ParseInt(os.Getenv("PORT"), 10, 16); err == nil {
		port = uint16(num64)
	} else {
		panic("Could not read port")
	}

	return &Env{
		DB:     conn,
		Port:   port,
		Host:   host,
		Logger: logger,
		DbCtx:  ctx,
	}
}

func (e *Env) GetPortString() string {
	return fmt.Sprintf(":%s", e.Port)
}

type responseObj struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (e *Env) WriteJsonResponse(w io.Writer, v any, msg string) {
	response := responseObj{Data: v}
	if msg != "" {
		response.Message = msg
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatalf("Error writing JSON response: %+v", response)
	}
}

type ControllerFunc func(e *Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error

// ! TO DEPRECATE ALL BELOW
// Use adapters/gateway.go instead
func (e *Env) Handle(fn ControllerFunc, protected bool) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		logger := e.Logger

		// CORS
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000/")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		)
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if protected && !isProtected(r) {
			logger.Error("Unauthenticated: Please log in to access this")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		if err := fn(e, w, r, p); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Content-Type-Options", "nosniff")

			switch err := err.(type) {
			case IApiError:
				logger.Error(err.Error())
				err.WriteJsonResponse(w)
			default:
				logger.Error(fmt.Sprintf("Internal error occurred: %v", err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				e.WriteJsonResponse(w, nil, "Uh oh... we need a minute")
			}
		}
	}
}

func isProtected(r *http.Request) bool {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return false
	}
	tokenString := cookie.Value
	if err := auth.VerifyToken(tokenString); err != nil {
		return false
	}
	return true
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
				err.WriteJsonResponse(w)
			default:
				logger.Error(fmt.Sprintf("Internal error occurred: %v", err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				e.WriteJsonResponse(w, nil, "Uh oh... we need a minute")
			}
		}
	}
}

type IApiError interface {
	Error() string
	GetStatus() int
	WriteJsonResponse(http.ResponseWriter)
}

type HttpStatusError struct {
	Message         string `json:"error"`
	DetailedMessage string
	Status          int    `json:"status"`
	Path            string `json:"path"`
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

func (err HttpStatusError) WriteJsonResponse(w http.ResponseWriter) {
	w.WriteHeader(err.Status)
	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(&err); e != nil {
		log.Fatalf("Error writing JSON response: %v", err)
	}
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
