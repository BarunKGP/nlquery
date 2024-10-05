-- +goose Up

CREATE TABLE IF NOT EXISTS users(
	id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	createdAt TIMESTAMPTZ NOT NULL default (timezone('utc', now())), 
	lastModified TIMESTAMPTZ NOT NULL default (timezone('utc', now())),
	name text NOT NULL,
	email VARCHAR(64) UNIQUE NOT NULL,
	providerUserId text,
	imageSrc text
);

-- +goose Down
DROP TABLE users;
