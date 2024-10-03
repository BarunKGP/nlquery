-- +goose Up

USE nlquery;

CREATE TABLE IF NOT EXISTS users(
	id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	createdAt TIMESTAMPTZ NOT NULL, 
	lastModified TIMESTAMPTZ NOT NULL,
	name text NOT NULL,
	email VARCHAR(64) UNIQUE NOT NULL,
	providerUserId text
	imageSrc text,
	sessionId bigint REFERENCES sessions(id)
);

CREATE TABLE IF NOT EXISTS sessions(
	id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	createdAt TIMESTAMPTZ NOT NULL,
	isAlive boolean default "true",
	killedAt TIMESTAMPTZ
	activeConnections int default 1
);

-- +goose Down
DROP TABLE users;
DROP TABLE sessions;
