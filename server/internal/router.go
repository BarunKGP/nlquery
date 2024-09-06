package internal

import (
	"github.com/julienschmidt/httprouter"
)

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

func GetApiRouter() *ApiRouter {
	return &ApiRouter{httprouter.New(), "/api/v1"}
}
