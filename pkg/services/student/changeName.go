package student

import (
	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) ChangeName(id int64, name string) (int64, *apperrors.Error) {
	id, err := s.Repo.ChangeName(id, name)

	if err != nil {
		return 0, apperrors.NewError(err)
	}

	return id, nil
}
