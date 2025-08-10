package subject

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) GetByCourseAndPeriod(course, period byte) ([]domain.Subject, *apperrors.Error) {
	subject, err := s.Repo.GetByCourseAndPeriod(course, period)

	if err != nil {
		return nil, apperrors.NewError(err)
	}

	return subject, nil
}
