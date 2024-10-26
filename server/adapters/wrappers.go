package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/BarunKGP/nlquery/core/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type ApiObject interface {
	ToJson() (string, error)
	FromJson(string) (ApiObject, error)
}

type ApiUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserId   string `json:"id"`
	ImageSrc string `json:"image"`
}

func NewApiUser() *ApiUser {
	return &ApiUser{}
}

func (a *ApiUser) ToJson() (string, error) {
	res, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (a *ApiUser) FromHttp(r *http.Request) (*ApiUser, error) {
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		return nil, fmt.Errorf("Unable to parse request body")
	}
	return a, nil
}

func (a *ApiUser) IsPersisted(ctx context.Context, db *database.Queries) bool {
	if _, err := db.GetUserByEmail(ctx, a.Email); err != nil {
		return false
	}
	return true
}

func (a *ApiUser) Persist(ctx context.Context, db *database.Queries) (database.User, error) {
	if ok := a.IsPersisted(ctx, db); ok {
		return db.GetUserByEmail(ctx, a.Email)

	}
	user, err := db.CreateUser(ctx, database.CreateUserParams{
		Name:           a.Name,
		Email:          a.Email,
		Imagesrc:       pgtype.Text{String: a.ImageSrc, Valid: true},
		Provideruserid: pgtype.Text{String: a.UserId, Valid: true},
	})
	if err != nil {
		slog.Error("Could not persist user", "user", *a, "error", err)
		return database.User{}, fmt.Errorf("Failed to persist user")
	}
	return user, nil
}
