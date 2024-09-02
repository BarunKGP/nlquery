package main

import (
	"log"
	"net/http"

	"github.com/BarunKGP/nlquery/controllers"
	"github.com/BarunKGP/nlquery/utils"
	"github.com/julienschmidt/httprouter"
)


type Signin struct {
	User    any    `json:user`
	Account string `json:account`
}

type User struct {
	Id       string `json: id`
	Name     string `json: name`
	Email    string `json: email`
	ImageSrc string `json: image`
}

type ApiRouter struct {
	*httprouter.Router
	prefix string
}

func (r *ApiRouter) Get(path string, handle httprouter.Handle) {
	r.GET(r.prefix + path, handle)
}

func (r *ApiRouter) Post(path string, handle httprouter.Handle) {
	r.POST(r.prefix + path, handle)
}

func (r *ApiRouter) Put(path string, handle httprouter.Handle) {
	r.PUT(r.prefix + path, handle)
}

func (r *ApiRouter) Patch(path string, handle httprouter.Handle) {
	r.PATCH(r.prefix + path, handle)
}

func (r *ApiRouter) Delete(path string, handle httprouter.Handle) {
	r.DELETE(r.prefix + path, handle)
}


func getRouter() *ApiRouter {
	return &ApiRouter{httprouter.New(), "/api/v1"}	
}

// func enableCors(next httprouter.Handle) httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 	next(w, r, p)
// 	}
// }

func main() {
	logger := utils.CreateLogger()
	router := getRouter()

	router.Get("/", controllers.HandleHome(logger))
	// router.Post("/auth/login", controllers.HandleSignin(logger))
	// router.Get("/me", controllers.GetCurrentUser(logger))

	logger.Info("Starting http server")
	log.Fatal(http.ListenAndServe(":8001", router))

}
