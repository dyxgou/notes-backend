package subject

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) GetByNameAndCourse(name string, course, period byte) (*domain.Subject, *apperrors.Error) {
	subject, err := s.Repo.GetByNameAndCourse(name, course, period)

	if err != nil {
		return nil, apperrors.NewError(err)
	}

	return subject, nil
}
