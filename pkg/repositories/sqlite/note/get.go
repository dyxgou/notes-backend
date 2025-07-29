package note

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) Get(
	id int64,
) (*domain.Note, error) {
	query := "SELECT * FROM note WHERE id = ?;"

	row := r.Db.QueryRow(query, id)

	var n domain.Note
	err := row.Scan(&n.Id, &n.GradeId, &n.StudentId, &n.Value)
	if err != nil {
		return nil, err
	}

	return &n, nil
}
