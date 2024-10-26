package core

import (
	"fmt"
	"net/http"

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

func NewApiRouter(prefix string) *ApiRouter {
	return &ApiRouter{Router: httprouter.New(), prefix: prefix}
}

// * NOTE: If this doesn't work, enable the CORS setting from adapters/gateway.go middleware()
func (a *ApiRouter) EnableCors(allowedOrigins []string) {
	a.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			origin := r.Header.Get("Origin")
			for _, ao := range allowedOrigins {
				if ao == origin {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Add("Access-Control-Allow-Credentials", "true")
					w.Header().Add(
						"Access-Control-Allow-Headers",
						"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
					)
					w.Header().Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
					break
				}
			}
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

var _ http.Handler = (*ApiRouter)(nil)

func (a *ApiRouter) HandleRoute(method string, endpoint string, handle httprouter.Handle) error {
	switch method {
	case http.MethodGet:
		a.Get(endpoint, handle)
	case http.MethodPost:
		a.Post(endpoint, handle)
	case http.MethodPut:
		a.Put(endpoint, handle)
	case http.MethodPatch:
		a.Patch(endpoint, handle)
	case http.MethodDelete:
		a.Delete(endpoint, handle)
	default:
		return fmt.Errorf("HTTP method %s not supported", method)
	}
	return nil
}
