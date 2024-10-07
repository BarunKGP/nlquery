-- +goose Up
CREATE TYPE thread_status AS enum('active', 'dormant', 'deleted');

CREATE TABLE IF NOT EXISTS threads(
	id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	createdAt TIMESTAMPTZ NOT NULL default (now() at time zone 'utc'), 
	lastModified TIMESTAMPTZ NOT NULL default (timezone('utc', now())),
	query text NOT NULL,
	userId bigint REFERENCES users(id),
	status thread_status default 'active',
	threadFileId bigint REFERENCES thread_files(id)
);

CREATE TABLE IF NOT EXISTS thread_files(
	id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	userId bigint REFERENCES users(id),
	createdAt TIMESTAMPTZ NOT NULL default (timezone('utc', now())),
	lastModified TIMESTAMPTZ NOT NULL default (timezone('utc', now())),
	columns text NOT NULL,
	types text NOT NULL
)


-- +goose Down
DROP TABLE threads, thread_files;
DROP TYPE thread_status;
