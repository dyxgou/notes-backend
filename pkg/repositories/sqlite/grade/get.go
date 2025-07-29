package grade

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Repository) Get(id int64) (*domain.Grade, error) {
	query := "SELECT * FROM grade WHERE id = ?;"

	row := r.Db.QueryRow(query, id)

	var g domain.Grade
	err := row.Scan(&g.Id, &g.Name, &g.SubjectId)
	if err != nil {
		return nil, err
	}

	return &g, nil
}
