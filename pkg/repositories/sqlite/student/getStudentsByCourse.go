package student

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) GetStudentsByCourse(courseId int64) ([]domain.Student, error) {
	query := "SELECT id, name FROM student WHERE course = ? ORDER BY name ASC;"

	rows, err := r.Db.Query(query, courseId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	students := make([]domain.Student, 0, 20)

	for rows.Next() {
		var s domain.Student

		err := rows.Scan(&s.Id, &s.Name)
		if err != nil {
			return nil, err
		}

		students = append(students, s)
	}

	return students, nil
}
