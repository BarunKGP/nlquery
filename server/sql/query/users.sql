-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
	name, email, provider, createdAt, lastModified
) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
	set name=$2,
	email=$3,
	lastModified=$4
WHERE id = $1;

-- name: UpdateProvider :exec
UPDATE users
	SET provider = $2
	WHERE id = $1;

