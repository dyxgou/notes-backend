-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS student (
  id INTEGER PRIMARY KEY NOT NULL,
  name VARCHAR(40) CHECK (LENGTH (name) <= 40),
  course SMALLINT CHECK (
    course >= 0
    AND course <= 11
  ),
  parent_phone VARCHAR(10) CHECK (
    LENGTH (parent_phone) = 10
    AND parent_phone NOT GLOB '*[^0-9]*'
  )
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE student;

-- +goose StatementEnd
