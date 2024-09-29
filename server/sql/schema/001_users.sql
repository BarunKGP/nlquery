-- +goose Up

CREATE TABLE IF NOT EXISTS users(
	id BIGSERIAL PRIMARY KEY,
	name text  NOT NULL,
	email text NOT NULL,
	providerUserId text,
	imageSrc text,
	createdAt timestamp NOT NULL, 
	lastModified timestamp NOT NULL
);

-- +goose Down
DROP TABLE users;
