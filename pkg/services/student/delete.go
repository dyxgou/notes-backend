package student

import (
	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) Delete(id int64) (int64, *apperrors.Error) {
	studentId, err := s.Repo.Delete(id)

	if err != nil {
		return 0, apperrors.NewError(err)
	}

	return studentId, nil
}
