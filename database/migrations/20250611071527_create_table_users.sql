-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   first_name VARCHAR(50) NOT NULL UNIQUE,
   last_name VARCHAR(16) NOT NULL,
   status VARCHAR(16) NOT NULL,
   email VARCHAR(16) NOT NULL,
   phone_number VARCHAR(16) NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
