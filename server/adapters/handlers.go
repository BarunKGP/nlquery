package adapters

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

	"github.com/BarunKGP/nlquery/core"
	"github.com/BarunKGP/nlquery/core/auth"
	"github.com/BarunKGP/nlquery/core/database"
	"github.com/BarunKGP/nlquery/ports"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

type Env struct {
	DB     ports.PersistentConn
	Port   uint16
	Host   string
	Logger *slog.Logger
	DbCtx  context.Context
}

func InitEnv() *Env {
	if e := godotenv.Load(); e != nil {
		log.Fatal("Unable to read environment variables")
	}
	dbUrl := getDbUrl()
	// TODO: Change level based on env: debug for dev, info for stg, warn for prod
	logger := core.CreateLogger(slog.LevelDebug)
	logger.Debug("Connecting to postgres db with dbUrl: " + dbUrl)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

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
	return fmt.Sprintf(":%d", e.Port)
}

func (e *Env) GetDbFactory() *database.Queries {
	return database.New(e.DB)
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

type responseObj struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func WriteJsonResponse(w io.Writer, v any, msg string) {
	response := responseObj{Data: v}
	if msg != "" {
		response.Message = msg
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatalf("Error writing JSON response: %+v", response)
	}
}

type ControllerFunc func(e *Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error
