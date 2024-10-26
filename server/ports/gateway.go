package ports

import (
	"net/http"

	"github.com/BarunKGP/nlquery/core/database"
	"github.com/julienschmidt/httprouter"
)

type RequestServer interface {
}

type Routeable interface {
	GetMethod() string
	GetEndpoint() string
	// GetController() ControllerFunc
}

type EnvReader interface {
	GetDbFactory() *database.Queries
}

type ControllerFunc func(EnvReader, http.ResponseWriter, *http.Request, httprouter.Params) error
