package subject

import (
	"database/sql"
	"errors"

	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) CreateAndGetSubject(subject *domain.Subject) (
	*domain.Subject, *apperrors.Error,
) {
	sub, err := s.Repo.GetByNameAndCourse(
		subject.Name, subject.Course, subject.Period,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			id, err := s.Repo.Insert(subject)
			if err != nil {
				return nil, apperrors.NewError(err)
			}

			sub, err := s.Repo.GetSubjectById(id)
			if err != nil {
				return nil, apperrors.NewError(err)
			}

			return sub, nil
		}

		return nil, apperrors.NewError(err)
	}

	return sub, nil
}
