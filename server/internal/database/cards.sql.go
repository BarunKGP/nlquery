// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: cards.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCard = `-- name: CreateCard :one
INSERT INTO cards(query, userId) 
VALUES ($1, $2)
RETURNING id, createdat, lastmodified, query, userid, status
`

type CreateCardParams struct {
	Query  pgtype.Text
	Userid pgtype.Int8
}

func (q *Queries) CreateCard(ctx context.Context, arg CreateCardParams) (Card, error) {
	row := q.db.QueryRow(ctx, createCard, arg.Query, arg.Userid)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Createdat,
		&i.Lastmodified,
		&i.Query,
		&i.Userid,
		&i.Status,
	)
	return i, err
}

const getCard = `-- name: GetCard :one
SELECT id, createdat, lastmodified, query, userid, status FROM cards
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCard(ctx context.Context, id int64) (Card, error) {
	row := q.db.QueryRow(ctx, getCard, id)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Createdat,
		&i.Lastmodified,
		&i.Query,
		&i.Userid,
		&i.Status,
	)
	return i, err
}

const getUserCards = `-- name: GetUserCards :many
SELECT (id, query, status, createdAt, lastModified)
FROM cards  
WHERE userId = $1
`

func (q *Queries) GetUserCards(ctx context.Context, userid pgtype.Int8) ([]interface{}, error) {
	rows, err := q.db.Query(ctx, getUserCards, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []interface{}
	for rows.Next() {
		var column_1 interface{}
		if err := rows.Scan(&column_1); err != nil {
			return nil, err
		}
		items = append(items, column_1)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCardStatus = `-- name: UpdateCardStatus :exec
UPDATE cards
SET status = $2, lastModified = (timezone('utc', now()))
WHERE id = $1
`

type UpdateCardStatusParams struct {
	ID     int64
	Status NullCardStatus
}

func (q *Queries) UpdateCardStatus(ctx context.Context, arg UpdateCardStatusParams) error {
	_, err := q.db.Exec(ctx, updateCardStatus, arg.ID, arg.Status)
	return err
}
