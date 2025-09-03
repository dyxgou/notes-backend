package note

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) GetByGradeIdAndStudentId(
	gradeId, studentId int64,
) (*domain.Note, error) {
	query := "SELECT id, value FROM note WHERE grade_id = ? AND student_id = ?;"

	row := r.Db.QueryRow(query, gradeId, studentId)

	var n domain.Note
	err := row.Scan(&n.Id, &n.Value)
	if err != nil {
		return nil, err
	}

	return &n, nil
}
