package subject

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) Create(subject *domain.Subject) (int64, *apperrors.Error) {
	id, err := s.Repo.Insert(subject)

	if err != nil {
		return 0, apperrors.NewError(err)
	}

	return id, nil
}
