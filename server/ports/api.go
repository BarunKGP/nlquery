package ports

import "net/http"

type ApiObject interface {
	ToJson() (string, error)
	FromJson(string) (ApiObject, error)
}

type ApiErrorable interface {
	Error() string
	GetStatus() int
	WriteJsonResponse(http.ResponseWriter)
}
