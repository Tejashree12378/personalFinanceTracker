-- +goose Up
-- +goose StatementBegin
CREATE TABLE limits (
      id SERIAL PRIMARY KEY,
      user_id integer NOT NULL,
      monthly_limit integer,
      yearly_limit integer,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      deleted_at TIMESTAMP,
      CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE limits;
-- +goose StatementEnd