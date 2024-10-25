package adapters

import (
	"fmt"
	"net/http"

	"github.com/BarunKGP/nlquery/core"
	"github.com/BarunKGP/nlquery/internal"
	"github.com/julienschmidt/httprouter"
)

type Gateway struct {
	// accepted origins for CORS
	AcceptedOrigins []string
	// prefix applied to all routes in the gateway
	ApiPrefix string

	// store for passed in env variables
	env *internal.Env
	// internal router for http requests
	router *internal.ApiRouter
}

func NewGateway(prefix string, origins []string) *Gateway {
	return &Gateway{AcceptedOrigins: origins, ApiPrefix: prefix}
}

func (g *Gateway) WithRouter(r *internal.ApiRouter) *Gateway {
	g.router = r
	return g
}

func (g *Gateway) WithEnv(e *internal.Env) *Gateway {
	g.env = e
	return g
}

func (g *Gateway) Init() error {
	// Initialize env vars
	g.env = internal.InitEnv()
	defer g.env.DB.Close(g.env.DbCtx)

	// Initialize router and routes
	g.router = internal.NewApiRouter(g.ApiPrefix)
	g.router.EnableCors(g.AcceptedOrigins)
	routes := InitRoutes()
	g.injectRoutes(routes)

	return http.ListenAndServe(g.env.GetPortString(), g.router)
}

func (g *Gateway) injectRoutes(routes map[*httpRoute]internal.ControllerFunc) {
	for r, cont := range routes {
		g.router.HandleRoute(r.GetMethod(), r.GetEndpoint(), g.handleRequest(cont))
	}
}

func (a *Gateway) handleRequest(fn internal.ControllerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := fn(a.env, w, r, p); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Content-Type-Options", "nosniff")

			switch err := err.(type) {
			case core.ApiErrorable:
				a.env.Logger.Error(err.Error())
				err.WriteJsonResponse(w)
			default:
				a.env.Logger.Error(fmt.Sprintf("Internal error occurred: %v", err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				a.env.WriteJsonResponse(w, nil, "Uh oh... we need a minute")
			}
		}

	}
}
