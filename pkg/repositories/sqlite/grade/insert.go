package grade

import (
	"database/sql"
	"fmt"

	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) Insert(g *domain.Grade) (int64, error) {
	tx, err := r.Db.Begin()
	if err != nil {
		return 0, err
	}

	q1 := "INSERT INTO grade(subject_id, name, is_final_exam) VALUES(?, ?, ?);"
	q2 := "UPDATE subject SET grades = grades + 1 WHERE id = ?;"
	q3 := "UPDATE subject SET has_final_exam = TRUE WHERE id = ?;"

	res, err := tx.Exec(q1, g.SubjectId, g.Name, g.IsFinalExam)
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

	if g.IsFinalExam {
		hasFinalExam, err := hasSubjectFinalExam(tx, g.SubjectId)
		if err != nil {
			return 0, err
		}

		if hasFinalExam {
			if err := tx.Rollback(); err != nil {
				return 0, fmt.Errorf("Rollback error: %w", err)
			}

			return 0, apperrors.ErrSubjectHasFinalExam
		}

		if _, err := tx.Exec(q3, g.SubjectId); err != nil {
			if err := tx.Rollback(); err != nil {
				return 0, fmt.Errorf("Rollback error: %w", err)
			}

			return 0, err
		}
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

func hasSubjectFinalExam(tx *sql.Tx, id int64) (bool, error) {
	q := "SELECT has_final_exam FROM subject WHERE id = ?;"

	var hasFinalGrade bool
	err := tx.QueryRow(q, id).Scan(&hasFinalGrade)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return hasFinalGrade, fmt.Errorf("Rollback error: %w", err)
		}
		return hasFinalGrade, err
	}

	return hasFinalGrade, nil
}
