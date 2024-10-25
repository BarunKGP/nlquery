package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteObjectToJson(w http.ResponseWriter, v any) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("Unable to convert payload: %v to json", v)
	}

	return nil
}

/**

{
	"error": "resource not found",
	"reason": "/auth path does not exist",
	"status": 404,
	"internal": {
		"errorCode": 4041
		"path": "",
		"fullSource": "",
		"message": ""
	}
}

**/

// func WriteErrorToJson(w http.ResponseWriter, msg string) {
//
// }
