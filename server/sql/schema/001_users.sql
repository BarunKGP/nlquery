-- +goose Up

CREATE TABLE IF NOT EXISTS users(
	id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	createdAt TIMESTAMPTZ NOT NULL, 
	lastModified TIMESTAMPTZ NOT NULL,
	name text NOT NULL,
	email VARCHAR(64) UNIQUE NOT NULL,
	providerUserId text,
	imageSrc text
);

-- +goose Down
DROP TABLE users;
