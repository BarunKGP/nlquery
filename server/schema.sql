CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL PRIMARY KEY,
	name text  NOT NULL,
	email text NOT NULL,
	provider text,
	createdAt timestamp NOT NULL, 
	lastModified timestamp NOT NULL
)
