package student

import (
	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) GetParentPhone(id int64) (string, *apperrors.Error) {
	tel, err := s.Repo.GetParentPhone(id)
	if err != nil {
		return tel, apperrors.NewError(err)
	}

	return tel, nil
}
