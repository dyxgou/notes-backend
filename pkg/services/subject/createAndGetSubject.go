package subject

import (
	"database/sql"
	"errors"

	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) CreateAndGetSubject(subject *domain.Subject) (
	int64, *apperrors.Error,
) {
	id, err := s.Repo.GetByNameAndCourse(
		subject.Name, subject.Course, subject.Period,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			id, err := s.Repo.Insert(subject)
			if err != nil {
				return 0, apperrors.NewError(err)
			}

			return id, nil
		}

		return 0, apperrors.NewError(err)
	}

	return id, nil
}
