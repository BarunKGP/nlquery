package controllers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/BarunKGP/nlquery/ports"
)

/* var _ Routeable = (*httpRoute)(nil)
 */

type httpRoute struct {
	method     string
	endpoint   string
	Controller ports.ControllerFunc
}

func NewHttpRoute(method string, endpoint string) ports.Routeable {
	return httpRoute{method: method, endpoint: endpoint}
}

func (h httpRoute) GetMethod() string {
	return h.method
}

func (h httpRoute) GetEndpoint() string {
	return h.endpoint
}

func InitRoutes() map[ports.Routeable]ports.ControllerFunc {

	t := map[string]ports.ControllerFunc{
		"GET /":     HandleHome,
		"GET /test": HandleTest,

		"POST /auth/signin": HandleSignin,
		"POST /auth/logout": (HandleLogout),

		"POST /user":    isProtected(HandleCreateUser),
		"GET /user/:id": isProtected(HandleGetUser),
	}

	res := make(map[ports.Routeable]ports.ControllerFunc)
	for rs, cf := range t {
		hr, err := ParseRoute(rs)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse route %s", rs))
		}
		res[hr] = cf
	}
	slog.Debug("parsed routes")
	return res
}

func ParseRoute(route string) (ports.Routeable, error) {
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

func (hr *httpRoute) AddController(fn ports.ControllerFunc) *httpRoute {
	hr.Controller = fn
	return hr
}
