-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS note (
  id INTEGER PRIMARY KEY NOT NULL,
  grade_id INT NOT NULL,
  student_id INT NOT NULL,
  value SMALLINT NOT NULL CHECK (
    value >= 10
    AND value <= 50
  ) DEFAULT 10,
  UNIQUE (grade_id, student_id),
  FOREIGN KEY (grade_id) REFERENCES grade (id) ON DELETE CASCADE,
  FOREIGN KEY (student_id) REFERENCES student (id) ON DELETE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE note;

-- +goose StatementEnd
