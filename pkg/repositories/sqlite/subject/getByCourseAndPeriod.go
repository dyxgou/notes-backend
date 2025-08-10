package subject

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) GetByCourseAndPeriod(course, period byte) ([]domain.Subject, error) {
	query := "SELECT id, name FROM subject WHERE course = ? AND period = ?;"

	subjects := make([]domain.Subject, 0, 14)

	rows, err := r.Db.Query(query, course, period)
	if err != nil {
		return subjects, err
	}

	defer rows.Close()

	for rows.Next() {
		var s domain.Subject

		err := rows.Scan(&s.Id, &s.Name)
		if err != nil {
			return subjects, err
		}

		subjects = append(subjects, s)
	}

	return subjects, nil
}
