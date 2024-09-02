package controllers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func HandleHome(logger *slog.Logger) httprouter.Handle {
    logger.Info("Home route")

    return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        fmt.Fprintf(w, "Hello from nlQuery\n")
    }
}

// func HandleSignin(logger *slog.Logger) httprouter.Handle {
//     return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		body, err := io.ReadAll(r.Body)
//         if err != nil {
// 			errMsg := "Unable to read body"
// 			logger.Error(errMsg, err)
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Header().Add("X-Error", errMsg)

// 			return
// 		}

		


// 	}
// }