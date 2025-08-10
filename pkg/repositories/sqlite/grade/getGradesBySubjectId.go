package grade

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) GetGradesBySubjectId(subjectId int64) ([]domain.Grade, error) {
	query := `
SELECT
  id,
  name,
  is_final_exam
FROM
  grade
WHERE
  subject_id = ?
ORDER BY
  is_final_exam ASC;`

	rows, err := r.Db.Query(query, subjectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	grades := make([]domain.Grade, 0, 10)

	for rows.Next() {
		var g domain.Grade

		if err := rows.Scan(&g.Id, &g.Name, &g.IsFinalExam); err != nil {
			return grades, err
		}

		grades = append(grades, g)
	}

	return grades, nil
}
