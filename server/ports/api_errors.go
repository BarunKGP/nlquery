package ports

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BarunKGP/nlquery/core"
)

type HttpStatusError struct {
	Message         string `json:"error"`
	DetailedMessage string
	Status          int    `json:"status"`
	Path            string `json:"path"`
}
func NewHttpError(errMsg string, status int, path string) HttpStatusError {
	// Build a new HttpStatusError
	// This can be written to a logger using `httpErr.Error()` or returned as an error value
	return HttpStatusError{
		Message: errMsg,
		Status:  status,
		Path:    path,
	}
}

func (err HttpStatusError) Error() string {
	return fmt.Sprintf("Error: %v at %v: returning HTTP %d",
		err.Message, err.Path, err.Status)
}

func (err HttpStatusError) GetStatus() int {
	return err.Status
}

func (err HttpStatusError) GetPath() string {
	return err.Path
}

func (err HttpStatusError) WriteJsonResponse(w http.ResponseWriter) {
	w.WriteHeader(err.Status)
	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(&err); e != nil {
		log.Fatalf("Error writing JSON response: %v", err)
	}
}

var _ core.ApiErrorable = (*HttpStatusError)(nil)
