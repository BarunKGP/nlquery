-- +goose Up

CREATE TABLE IF NOT EXISTS users(
	id BIGSERIAL PRIMARY KEY,
	createdAt TIMESTAMPTZ NOT NULL, 
	lastModified TIMESTAMPTZ NOT NULL,
	name VARCHAR(200)  NOT NULL,
	email VARCHAR(200) UNIQUE NOT NULL,
	providerUserId VARCHAR(64),
	imageSrc text,
	sessionId UUID,
	accountId UUID
);

-- +goose Down
DROP TABLE users;
