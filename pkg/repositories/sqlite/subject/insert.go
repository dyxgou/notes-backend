package subject

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) Insert(subject *domain.Subject) (int64, error) {
	query := "INSERT INTO subject(name, course, period) VALUES(?, ?, ?);"

	res, err := r.Db.Exec(query, subject.Name, subject.Course, subject.Period)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
