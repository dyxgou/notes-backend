package student

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) Get(id int64) (*domain.Student, error) {
	query := "SELECT * FROM student WHERE id = ?;"

	row := r.Db.QueryRow(query, id)

	var s domain.Student
	err := row.Scan(&s.Id, &s.Name, &s.Course, &s.ParentPhone)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
