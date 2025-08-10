package subject

import "github.com/dyxgou/notas/pkg/domain"

func (r *Repository) GetByNameAndCourse(name string, course, period byte) (*domain.Subject, error) {
	query := "SELECT * FROM subject WHERE name = ? AND course = ? AND period = ?;"

	row := r.Db.QueryRow(query, name, course, period)

	var s domain.Subject

	err := row.Scan(&s.Id, &s.Name, &s.Course, &s.Period, &s.HasFinalExam, &s.Grades)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
