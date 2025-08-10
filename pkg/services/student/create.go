package student

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) Create(student *domain.Student) (int64, *apperrors.Error) {
	id, err := s.Repo.Insert(student)
	if err != nil {
		return 0, apperrors.NewError(err)
	}

	return id, nil
}
