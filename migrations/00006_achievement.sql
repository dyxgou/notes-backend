-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS achievement (
  id INTEGER PRIMARY KEY NOT NULL,
  description VARCHAR(200) CHECK (LENGTH (description) <= 200),
  subject_name VARCHAR(15) CHECK (LENGTH (subject_name) <= 15),
  performance VARCHAR(7) NOT NULL CHECK (
    performance IN ('low', 'basic', 'high', 'superior')
  ),
  course SMALLINT CHECK (
    course >= 0
    AND course <= 11
  ),
  UNIQUE (subject_name, performance, course)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE achievement;

-- +goose StatementEnd
