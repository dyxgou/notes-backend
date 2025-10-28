-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS grade (
  id INTEGER PRIMARY KEY NOT NULL,
  name VARCHAR(20) CHECK (LENGTH (name) <= 20),
  subject_id INTEGER NOT NULL,
  is_final_exam BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (subject_id) REFERENCES subject (id)
);

-- Performance optimization: Add index on subject_id
CREATE INDEX IF NOT EXISTS idx_grade_subject_id ON grade (subject_id);

CREATE TRIGGER IF NOT EXISTS trg_check_final_grade_requirement BEFORE INSERT ON grade FOR EACH ROW WHEN NEW.is_final_exam = FALSE BEGIN
SELECT
  RAISE (
    FAIL,
    'Cannot add non-final grade: subject has reached maximum grades (9). Next grade must be final.'
  )
WHERE
  (
    SELECT
      COUNT(*)
    FROM
      grade
    WHERE
      subject_id = NEW.subject_id
      AND is_final_exam = FALSE
  ) = 9;

END;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE grade;

-- +goose StatementEnd
