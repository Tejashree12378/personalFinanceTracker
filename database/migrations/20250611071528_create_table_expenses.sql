-- +goose Up
-- +goose StatementBegin
CREATE TABLE expenses (
      id SERIAL PRIMARY KEY,
      user_id INTEGER NOT NULL,
      amount DECIMAL(10, 2) NOT NULL,
      description VARCHAR(255),
      name VARCHAR(255),
      category VARCHAR(100),
      date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      deleted_at TIMESTAMP,
      CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE expenses;
-- +goose StatementEnd