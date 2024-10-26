package controllers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/BarunKGP/nlquery/core/auth"
	"github.com/BarunKGP/nlquery/ports"
	"github.com/julienschmidt/httprouter"
)

type Middleware func(ports.ControllerFunc) ports.ControllerFunc

func isProtected(fn ports.ControllerFunc) ports.ControllerFunc {
	return func(e ports.EnvReader, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			slog.Debug("Unauthorized request on protected resource", "error", err)
			return fmt.Errorf("Please log in to continue")
		}
		tokenString := cookie.Value
		if err := auth.VerifyToken(tokenString); err != nil {
			slog.Debug("Unable to verify token", "error", err)
			return fmt.Errorf("Please log in to continue")
		}
		return fn(e, w, r, p)
	}
}
