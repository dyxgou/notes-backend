package grade

import (
	"log/slog"

	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) Delete(gradeId, subjectId int64) (int64, *apperrors.Error) {
	id, err := s.Repo.Delete(gradeId, subjectId)
	if err != nil {
		slog.Error("deleting grade", "err", err)
		return 0, apperrors.NewError(err)
	}

	return id, nil
}
