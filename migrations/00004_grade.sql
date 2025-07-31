-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS grade (
  id INTEGER PRIMARY KEY NOT NULL,
  name VARCHAR(20) CHECK (LENGTH (name) <= 20),
  subject_id INTEGER NOT NULL,
  FOREIGN KEY (subject_id) REFERENCES subject (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE grade;

-- +goose StatementEnd
