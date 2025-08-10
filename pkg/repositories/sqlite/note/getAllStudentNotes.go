package note

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) GetAllStudentNotes(
	studentId, subjectId int64,
) ([]domain.Note, error) {
	query := `
SELECT
  n.*
FROM
  grade g
  JOIN note n ON n.grade_id = g.id
  AND n.student_id = ?
WHERE
  subject_id = ?
ORDER BY
  is_final_exam ASC;`

	notes := make([]domain.Note, 0, 10)

	row, err := r.Db.Query(query, studentId, subjectId)
	if err != nil {
		return notes, err
	}

	for row.Next() {
		var n domain.Note

		err := row.Scan(&n.Id, &n.GradeId, &n.StudentId, &n.Value)
		if err != nil {
			return notes, err
		}

		notes = append(notes, n)
	}

	return notes, nil
}
