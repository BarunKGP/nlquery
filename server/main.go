package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BarunKGP/nlquery/controllers"
	"github.com/BarunKGP/nlquery/internal"
	"github.com/BarunKGP/nlquery/internal/auth"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// func middleware(next httprouter.Handle, logger *slog.Logger) httprouter.Handle {
// 	// Enable CORS
// 	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		// Logging
// 		logger.Info("Home route")
//
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		next(w, r, p)
// 	}
// }

var env *internal.Env

func init() {
	if e := godotenv.Load(); e != nil {
		log.Fatal("Unable to read environment variables")
	}

	logger := internal.CreateLogger()

	// Get db
	dbUrl := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	logger.Info("Connecting to postgres db with dbUrl: " + dbUrl)

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

	authConfig := auth.NewAuthConfig([]string{"google", "github"})

	env = &internal.Env{
		DB:         conn,
		Port:       os.Getenv("PORT"),
		Host:       host,
		Logger:     logger,
		DbCtx:      ctx,
		AuthConfig: authConfig,
	}

}

func main() {
	router := internal.GetApiRouter()
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "localhost:3000")
		}
		w.WriteHeader(http.StatusNoContent)
	})

	router.Get("/", env.Handle(controllers.HandleHome))
	router.Get("/user/:id", env.Handle(controllers.HandleGetUser))
	router.Post("/user", env.Handle(controllers.HandleCreateUser))

	// Auth
	router.Get("/auth", env.Handle(controllers.GetAuthProviders))
	// router.Get("/auth/:provider", env.Handle(controllers.HandleAuth))
	// router.Get("/auth/:provider/callback", env.Handle(controllers.CompleteAuth))
	// router.Get("/auth/:provider/callback", env.Handle(controllers.CompleteAuth))

	// router.Get("/", internal.Handle(internal.Handler{env, controllers.HandleHome}))

	// router.Post("/auth/logout", middleware(controllers.HandleSignout, logger))
	// router.Post("/auth/signup", middleware(controllers.HandleSignup, logger))

	// router.Patch("/user/:id", internal.Handle(internal.Handler{Env: env, H: controllers.HandleGetUser}))

	env.Logger.Info("Starting http server")
	log.Fatal(http.ListenAndServe(":8001", router))

}
