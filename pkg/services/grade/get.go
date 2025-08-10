package grade

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) Get(id int64) (*domain.Grade, *apperrors.Error) {
	grade, err := s.Repo.Get(id)

	if err != nil {
		return nil, apperrors.NewError(err)
	}

	return grade, nil
}
