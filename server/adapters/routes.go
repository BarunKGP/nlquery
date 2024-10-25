package adapters

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/BarunKGP/nlquery/controllers"
	"github.com/BarunKGP/nlquery/internal"
	"github.com/BarunKGP/nlquery/internal/auth"
	"github.com/julienschmidt/httprouter"
)

/*
	 type Routeable interface {
		GetMethod() string
		GetEndpoint() string
	}

var _ Routeable = (*httpRoute)(nil)
*/
type httpRoute struct {
	method     string
	endpoint   string
	Controller internal.ControllerFunc
}

func NewHttpRoute(method string, endpoint string) *httpRoute {
	return &httpRoute{method: method, endpoint: endpoint}
}

func (h httpRoute) GetMethod() string {
	return h.method
}

func (h httpRoute) GetEndpoint() string {
	return h.endpoint
}

type Middleware func(internal.ControllerFunc) internal.ControllerFunc

func isProtected(fn internal.ControllerFunc) internal.ControllerFunc {
	return func(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			e.Logger.Debug("Unauthorized request on protected resource", "error", err)
			return fmt.Errorf("Please log in to continue")
		}
		tokenString := cookie.Value
		if err := auth.VerifyToken(tokenString); err != nil {
			e.Logger.Debug("Unable to verify token", "error", err)
			return fmt.Errorf("Please log in to continue")
		}
		return fn(e, w, r, p)
	}
}

func ParseRoute(route string) (*httpRoute, error) {
	parts := strings.Split(route, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Improper formatting: '%s' should contain HTTP method and endpoint")
	}

	flag := false
	for _, m := range []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch} {
		if m == parts[0] {
			flag = true
			break
		}
	}
	if !flag {
		return nil, fmt.Errorf("Invalid HTTP method in %s", route)
	}

	method := parts[0]
	var endpoint string
	tmpl := "/%s"
	if parts[1][0] == '/' {
		endpoint = fmt.Sprintf(tmpl, parts[1][1:])
	} else {
		endpoint = fmt.Sprintf(tmpl, parts[1])
	}

	return NewHttpRoute(method, endpoint), nil
}

func (hr *httpRoute) AddController(fn internal.ControllerFunc) *httpRoute {
	hr.Controller = fn
	return hr
}

func InitRoutes() map[*httpRoute]internal.ControllerFunc {
	r := map[string]internal.ControllerFunc{
		"GET /":     controllers.HandleHome,
		"GET /test": controllers.HandleTest,

		"POST /auth/signin": controllers.HandleSignin,
		"POST /auth/logout": isProtected(controllers.HandleLogout),

		"POST /user":    isProtected(controllers.HandleCreateUser),
		"GET /user/:id": isProtected(controllers.HandleGetUser),
	}

	res := make(map[*httpRoute]internal.ControllerFunc)
	for rs, cf := range r {
		hr, err := ParseRoute(rs)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse route %s", rs))
		}
		res[hr] = cf
	}
	slog.Debug("parsed routes")
	return res
}

// func parseRoutes(method string, route string, controller internal.ControllerFunc) {
//
// }
