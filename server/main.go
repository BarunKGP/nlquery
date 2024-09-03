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

func getRouter() *ApiRouter {
	return &ApiRouter{httprouter.New(), "/api/v1"}
}

func middleware(next httprouter.Handle, logger *slog.Logger) httprouter.Handle {
	// Logging
	logger.Info("Home route")

	// Enable CORS
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		next(w, r, p)
	}
}

func main() {
	logger := utils.CreateLogger()
	router := getRouter()

	router.Get("/", middleware(controllers.HandleHome, logger))
	router.Post("/auth/login", middleware(controllers.HandleSignin, logger))

	logger.Info("Starting http server")
	log.Fatal(http.ListenAndServe(":8001", router))

}
