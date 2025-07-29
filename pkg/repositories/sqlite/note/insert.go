package note

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) Insert(note *domain.Note) (int64, error) {
	query := "INSERT INTO note(grade_id, student_id) VALUES(?, ?);"

	res, err := r.Db.Exec(query, note.GradeId, note.StudentId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
