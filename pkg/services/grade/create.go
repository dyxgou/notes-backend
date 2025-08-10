package grade

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) Create(grade *domain.Grade) (int64, *apperrors.Error) {
	id, err := s.Repo.Insert(grade)
	if err != nil {
		return 0, apperrors.NewError(err)
	}

	return id, nil
}
