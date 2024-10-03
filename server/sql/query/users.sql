-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
	name, email, providerUserId, imageSrc, createdAt, lastModified
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
	set name=$2,
	email=$3,
	lastModified=$4
WHERE id = $1;

-- name: UpdateProviderUserId :exec
UPDATE users
	SET providerUserId = $2
	WHERE id = $1;

