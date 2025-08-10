package student

import (
	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) GetStudentAverage(
	studentId, subjectId int64,
) (float64, *apperrors.Error) {
	avg, err := s.Repo.GetStudentAverage(studentId, subjectId)
	if err != nil {
		return 0.0, apperrors.NewError(err)
	}

	return avg, nil
}
