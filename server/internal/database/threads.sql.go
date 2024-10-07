// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: threads.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createThread = `-- name: CreateThread :one
INSERT INTO threads(userId, query, threadFileId) 
VALUES ($1, $2, $3)
RETURNING id, createdat, lastmodified, query, userid, status, threadfileid
`

type CreateThreadParams struct {
	Userid       pgtype.Int8
	Query        string
	Threadfileid pgtype.Int8
}

func (q *Queries) CreateThread(ctx context.Context, arg CreateThreadParams) (Thread, error) {
	row := q.db.QueryRow(ctx, createThread, arg.Userid, arg.Query, arg.Threadfileid)
	var i Thread
	err := row.Scan(
		&i.ID,
		&i.Createdat,
		&i.Lastmodified,
		&i.Query,
		&i.Userid,
		&i.Status,
		&i.Threadfileid,
	)
	return i, err
}

const createThreadFile = `-- name: CreateThreadFile :one
INSERT INTO thread_files(userId, columns, types) 
VALUES ($1, $2, $3)
RETURNING id
`

type CreateThreadFileParams struct {
	Userid  pgtype.Int8
	Columns string
	Types   string
}

func (q *Queries) CreateThreadFile(ctx context.Context, arg CreateThreadFileParams) (int64, error) {
	row := q.db.QueryRow(ctx, createThreadFile, arg.Userid, arg.Columns, arg.Types)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getThread = `-- name: GetThread :one
SELECT id, createdat, lastmodified, query, userid, status, threadfileid FROM threads
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetThread(ctx context.Context, id int64) (Thread, error) {
	row := q.db.QueryRow(ctx, getThread, id)
	var i Thread
	err := row.Scan(
		&i.ID,
		&i.Createdat,
		&i.Lastmodified,
		&i.Query,
		&i.Userid,
		&i.Status,
		&i.Threadfileid,
	)
	return i, err
}

const getUserThreads = `-- name: GetUserThreads :many
SELECT (
	id, 
	query, 
	(
		SELECT columns AS parsedColumns, types AS parsedTypes
		FROM thread_files
		WHERE t.threadFileId = id
		LIMIT 1

	), 
	createdAt, 
	lastModified
)
FROM threads AS t
WHERE t.userId = $1 AND t.status = 'active'
ORDER BY createdAt ASC
`

func (q *Queries) GetUserThreads(ctx context.Context, userid pgtype.Int8) ([]interface{}, error) {
	rows, err := q.db.Query(ctx, getUserThreads, userid)
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

const updateThreadStatus = `-- name: UpdateThreadStatus :exec
UPDATE threads
SET status = $2, lastModified = (timezone('utc', now()))
WHERE id = $1
`

type UpdateThreadStatusParams struct {
	ID     int64
	Status NullThreadStatus
}

func (q *Queries) UpdateThreadStatus(ctx context.Context, arg UpdateThreadStatusParams) error {
	_, err := q.db.Exec(ctx, updateThreadStatus, arg.ID, arg.Status)
	return err
}
