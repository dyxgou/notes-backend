package grade

import "github.com/dyxgou/notas/pkg/domain"

func (r *Repository) GetGradesBySubjectId(subjectId int64) ([]domain.Grade, error) {
	query := "SELECT id, name FROM grade WHERE subject_id = ?;"

	rows, err := r.Db.Query(query, subjectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	grades := make([]domain.Grade, 0, 10)

	for rows.Next() {
		var g domain.Grade

		if err := rows.Scan(&g.Id, &g.Name); err != nil {
			return grades, err
		}

		grades = append(grades, g)
	}

	return grades, nil
}
