-- +goose Up
CREATE TYPE card_status AS enum('active', 'dormant', 'deleted');

CREATE TABLE IF NOT EXISTS cards(
	id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	createdAt TIMESTAMPTZ NOT NULL default (now() at time zone 'utc'), 
	lastModified TIMESTAMPTZ NOT NULL default (timezone('utc', now())),
	query text,
	userId bigint REFERENCES users(id),
	status card_status default 'active'
);



-- +goose Down
DROP TABLE cards;
DROP TYPE card_status;
