package subject

import "github.com/dyxgou/notas/pkg/domain"

func (r *Repository) GetSubjectById(id int64) (*domain.Subject, error) {
	q := "SELECT * FROM subject WHERE id = ?;"

	var s domain.Subject
	err := r.Db.QueryRow(q, id).Scan(
		&s.Id, &s.Name, &s.Course, &s.Period, &s.HasFinalExam, &s.Grades,
	)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
