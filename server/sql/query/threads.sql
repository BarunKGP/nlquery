-- name: GetThread :one
SELECT * FROM threads
WHERE id = $1
LIMIT 1;

-- name: CreateThread :one
INSERT INTO threads(userId, query, threadFileId) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateThreadFile :one
INSERT INTO thread_files(userId, columns, types) 
VALUES ($1, $2, $3)
RETURNING id; 

-- name: GetUserThreads :many
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
ORDER BY createdAt ASC;

-- name: UpdateThreadStatus :exec
UPDATE threads
SET status = $2, lastModified = (timezone('utc', now()))
WHERE id = $1;



