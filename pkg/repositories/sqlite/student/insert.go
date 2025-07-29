package student

import (
	"context"
	"time"

	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) Insert(student *domain.Student) (int64, error) {
	query := "INSERT INTO student(name, course, parent_phone) VALUES(?, ?, ?);"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := r.Db.ExecContext(ctx, query, student.Name, student.Course, student.ParentPhone)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
