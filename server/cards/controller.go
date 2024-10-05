package cards

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BarunKGP/nlquery/internal"
	"github.com/BarunKGP/nlquery/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/julienschmidt/httprouter"
)

func HandleCreateCard(e *internal.Env, w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	queries := database.New(e.DB)

	iquery := apiCards{}
	if err := json.NewDecoder(r.Body).Decode(&iquery); err != nil {
		errMsg := fmt.Sprintf("Error decoding body: %v", err.Error())
		httpErr := internal.NewHttpError(errMsg, http.StatusInternalServerError, r.URL.Path)
		e.Logger.Error(httpErr.Error())
		return httpErr
	}

	e.Logger.Info(fmt.Sprintf("Received query from request body: %v", iquery))

	card, err := queries.CreateCard(
		e.DbCtx,
		database.CreateCardParams{
			Query: pgtype.Text{String: iquery.query, Valid: true}, Userid: pgtype.Int8{Int64: iquery.userId, Valid: true},
		})
	if err != nil {
		errMsg := fmt.Sprintf("Unable to create card: %v", err.Error())
		httpErr := internal.HttpStatusError{
			Message: errMsg,
			Status:  http.StatusInternalServerError,
			Path:    r.URL.Path,
		}
		e.Logger.Error(httpErr.Error())
	}

	e.Logger.Info(fmt.Sprintf("Card %d created successfully", card.ID))

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	e.WriteJsonResponse(w, card, "Card created successfully")

	return nil
}
