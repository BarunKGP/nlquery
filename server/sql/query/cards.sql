-- name: GetCard :one
SELECT * FROM cards
WHERE id = $1
LIMIT 1;

-- name: CreateCard :one
INSERT INTO cards(query, userId) 
VALUES ($1, $2)
RETURNING *;

-- name: GetUserCards :many
SELECT (id, query, status, createdAt, lastModified)
FROM cards  
WHERE userId = $1;

-- name: UpdateCardStatus :exec
UPDATE cards
SET status = $2, lastModified = (timezone('utc', now()))
WHERE id = $1;



