package core

import (
	"net/http"
)

type ApiErrorable interface {
	Error() string
	GetStatus() int
	WriteJsonResponse(http.ResponseWriter)
}
