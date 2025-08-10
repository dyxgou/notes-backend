package grade

import (
	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) ChangeName(id int64, name string) *apperrors.Error {
	err := s.Repo.ChangeName(id, name)

	if err != nil {
		return apperrors.NewError(err)
	}

	return nil
}
