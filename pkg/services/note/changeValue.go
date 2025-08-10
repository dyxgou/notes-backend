package note

import (
	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) ChangeValue(id int64, value byte) *apperrors.Error {
	err := s.Repo.ChangeValue(id, value)

	if err != nil {
		return apperrors.NewError(err)
	}

	return nil
}
