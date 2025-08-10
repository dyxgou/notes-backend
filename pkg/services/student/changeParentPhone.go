package student

import (
	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) ChangeParentPhone(id int64, phone string) (int64, *apperrors.Error) {
	id, err := s.Repo.ChangeParentPhone(id, phone)

	if err != nil {
		return 0, apperrors.NewError(err)
	}

	return id, nil
}
