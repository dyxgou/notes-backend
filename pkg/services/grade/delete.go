package grade

import (
	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) Delete(gradeId, subjectId int64) (int64, *apperrors.Error) {
	id, err := s.Repo.Delete(gradeId, subjectId)
	if err != nil {
		return 0, apperrors.NewError(err)
	}

	return id, nil
}
