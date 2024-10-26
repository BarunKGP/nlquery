package threads

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/BarunKGP/nlquery/adapters"
	"github.com/BarunKGP/nlquery/core/database"
	"github.com/BarunKGP/nlquery/ports"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/julienschmidt/httprouter"
)

type apiRequestFile struct {
	Columns []string `json:"columns"`
	Types   []string `json:"types"`
}

type apiRequestObj struct {
	Query string `json:"query"`
	//! We will have to get the user's identity separately
	// Use the `auth_token` from the cookie we created
	Id           int64           `json:"id"`
	FileDetails  *apiRequestFile `json:"file,omitempty"`
	ThreadFileId int64           `json:"threadFileId"`
}

func parseFileDetails(fileobj []string, sep string) (string, error) {
	var columnData string
	for _, cols := range fileobj {
		if columnData == "" {
			columnData = cols
			continue
		}
		columnData = fmt.Sprint("%s%s%s", columnData, sep, cols)
	}

	if len(strings.Split(columnData, sep)) != len(fileobj) {
		return "", fmt.Errorf(
			"Error creating `thread_file`: length of parsed columns: %d does not match length in request body: %d",
			len(fileobj),
			len(strings.Split(columnData, ";")),
		)
	}

	return columnData, nil
}

func HandleCreateThread(e ports.EnvReader, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	// Ensure user is logged in
	// Should this be handled in the Protected handler?
	user, err := adapters.NewApiUser().FromHttp(r)
	if err != nil {
		errMsg := fmt.Sprintf("Error decoding body: %v", err.Error())
		httpErr := adapters.NewHttpError(errMsg, http.StatusInternalServerError, r.URL.Path)
		slog.Error(httpErr.Error())
		return httpErr
	}
	slog.Info(fmt.Sprintf("Received query from request body: %v", jsonBody))

	dbFactory := e.GetDbFactory()
	var ctfParams database.CreateThreadParams
	//* NOTE: If user provides a new file or alters file schema in the
	// middle of a thread, it will result in a new threadFile being
	// created. Essentially, only one of the branches work at a time
	if jsonBody.FileDetails == nil {
		// We are adding a new query to an existing thread
		if jsonBody.ThreadFileId == 0 {
			//* We start db IDs from 1
			httpErr := adapters.NewHttpError("Invalid threadFieldId", http.StatusBadRequest, r.URL.Path)
			slog.Error(httpErr.Error())
			return httpErr
		}

		ctfParams = database.CreateThreadParams{
			Userid:       pgtype.Int8{Int64: jsonBody.Id, Valid: true},
			Threadfileid: pgtype.Int8{Int64: jsonBody.ThreadFileId, Valid: true},
			Query:        jsonBody.Query,
		}
	} else {
		// We need to create the `thread_file` first
		columnTypes, err := parseFileDetails(jsonBody.FileDetails.Types, ";")
		if err != nil {
			return fmt.Errorf("Unable to parse file details: %v", err)
		}
		columnNames, err := parseFileDetails(jsonBody.FileDetails.Columns, ";")
		if err != nil {
			return fmt.Errorf("Unable to parse file details: %v", err)
		}

		threadFileId, err := dbFactory.CreateThreadFile(
			r.Context(),
			database.CreateThreadFileParams{
				Columns: columnNames,
				Types:   columnTypes,
			},
		)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to create thread file: %v", err.Error())
			httpErr := adapters.HttpStatusError{
				Message: errMsg,
				Status:  http.StatusInternalServerError,
				Path:    r.URL.Path,
			}
			slog.Error(httpErr.Error())
			return httpErr
		}

		ctfParams = database.CreateThreadParams{
			Userid:       pgtype.Int8{Int64: jsonBody.Id, Valid: true},
			Threadfileid: pgtype.Int8{Int64: threadFileId, Valid: true},
			Query:        jsonBody.Query,
		}
	}

	thread, err := dbFactory.CreateThread(r.Context(), ctfParams)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to create thread file: %v", err.Error())
		httpErr := adapters.HttpStatusError{
			Message: errMsg,
			Status:  http.StatusInternalServerError,
			Path:    r.URL.Path,
		}
		slog.Error(httpErr.Error())
		return httpErr
	}
	slog.Info(fmt.Sprintf("thread %d created successfully", thread.ID))

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	adapters.WriteJsonResponse(
		w, map[string]string{"threadId": string(thread.ID)}, "Thread created successfully",
	)
	return nil
}
