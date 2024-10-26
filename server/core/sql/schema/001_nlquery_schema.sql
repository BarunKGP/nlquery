-- +goose Up
CREATE TABLE IF NOT EXISTS users(
	id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	createdAt TIMESTAMPTZ NOT NULL default (timezone('utc', now())), 
	lastModified TIMESTAMPTZ NOT NULL default (timezone('utc', now())),
	name text NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	providerUserId text,
	imageSrc text
); 

CREATE TYPE thread_status AS enum('active', 'dormant', 'deleted');

CREATE TABLE IF NOT EXISTS thread_files(
	id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	userId bigint REFERENCES users(id) ON DELETE CASCADE,
	createdAt TIMESTAMPTZ NOT NULL default (timezone('utc', now())),
	lastModified TIMESTAMPTZ NOT NULL default (timezone('utc', now())),
	columns text NOT NULL,
	types text NOT NULL
);

CREATE TABLE IF NOT EXISTS threads(
	id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	createdAt TIMESTAMPTZ NOT NULL default (now() at time zone 'utc'), 
	lastModified TIMESTAMPTZ NOT NULL default (timezone('utc', now())),
	query text NOT NULL,
	status thread_status default 'active',
	threadFileId bigint REFERENCES thread_files(id),
	userId bigint REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE threads;
DROP TABLE thread_files;
DROP TYPE thread_status;
DROP TABLE users;
