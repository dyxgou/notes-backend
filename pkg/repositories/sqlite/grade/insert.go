package grade

import (
	"fmt"

	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) Insert(g *domain.Grade) (int64, error) {
	tx, err := r.Db.Begin()
	if err != nil {
		return 0, err
	}

	q1 := "INSERT INTO grade(subject_id, name) VALUES(?, ?);"
	q2 := "UPDATE subject SET grades = grades + 1 WHERE id = ?;"
	res, err := tx.Exec(q1, g.SubjectId, g.Name)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, fmt.Errorf("Rollback error: %w", err)
		}
		return 0, err
	}

	_, err = tx.Exec(q2, g.SubjectId)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, fmt.Errorf("Rollback error: %w", err)
		}
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
