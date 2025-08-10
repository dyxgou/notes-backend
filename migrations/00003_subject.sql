-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS subject (
  id INTEGER PRIMARY KEY NOT NULL,
  name VARCHAR(15) CHECK (LENGTH (name) <= 15),
  course SMALLINT CHECK (
    course >= 0
    AND course <= 11
  ),
  period SMALLINT CHECK (
    period >= 1
    AND period <= 4
  ),
  has_final_exam BOOLEAN DEFAULT FALSE,
  grades SMALLINT CHECK (
    grades >= 0
    AND grades <= 10
  ) DEFAULT 0,
  UNIQUE (course, period, name)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE subject;

-- +goose StatementEnd
