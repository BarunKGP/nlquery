package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/BarunKGP/nlquery/controllers"
	"github.com/BarunKGP/nlquery/utils"
	"github.com/julienschmidt/httprouter"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageSrc string `json:"image"`
	Password string `json:"-"`
}

type ApiRouter struct {
	*httprouter.Router
	prefix string
}

func (r *ApiRouter) Get(path string, handle httprouter.Handle) {
	r.GET(r.prefix+path, handle)
}

func (r *ApiRouter) Post(path string, handle httprouter.Handle) {
	r.POST(r.prefix+path, handle)
}

func (r *ApiRouter) Put(path string, handle httprouter.Handle) {
	r.PUT(r.prefix+path, handle)
}

func (r *ApiRouter) Patch(path string, handle httprouter.Handle) {
	r.PATCH(r.prefix+path, handle)
}

func (r *ApiRouter) Delete(path string, handle httprouter.Handle) {
	r.DELETE(r.prefix+path, handle)
}

func getApiRouter() *ApiRouter {
	return &ApiRouter{httprouter.New(), "/api/v1"}
}

func middleware(next httprouter.Handle, logger *slog.Logger) httprouter.Handle {
	// Enable CORS
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Logging
		logger.Info("Home route")

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		next(w, r, p)
	}
}

type apiHandler func(http.ResponseWriter, *http.Request) error

type ErrHandle func(http.ResponseWriter, *http.Request, httprouter.Params) error

func handleError(eh ErrHandle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := eh(w, r, p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func main() {
	logger := utils.CreateLogger()
	router := getApiRouter()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "localhost:3000")
		}
		w.WriteHeader(http.StatusNoContent)
	})

	router.Get("/", handleError(controllers.HandleHome))
	router.Post("/auth/login", handleError(controllers.HandleSignin))
	// router.Post("/auth/logout", middleware(controllers.HandleSignout, logger))
	// router.Post("/auth/signup", middleware(controllers.HandleSignup, logger))

	logger.Info("Starting http server")
	log.Fatal(http.ListenAndServe(":8001", router))

}
