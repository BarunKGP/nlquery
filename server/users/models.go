// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package users

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID       int64
	Name     string
	Email    string
	Provider pgtype.Text
}
