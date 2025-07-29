package note

import (
	"database/sql"
	"errors"

	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) CreateAndGet(note *domain.Note) (*domain.Note, *domain.AppError) {
	n, err := s.Repo.GetByGradeIdAndStudentId(note.GradeId, note.StudentId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			id, err := s.Repo.Insert(note)
			if err != nil {
				return nil, domain.NewError(err)
			}

			n, err := s.Repo.Get(id)
			if err != nil {
				return nil, domain.NewError(err)
			}

			return n, nil
		}

		return nil, domain.NewError(err)
	}

	return n, nil
}
