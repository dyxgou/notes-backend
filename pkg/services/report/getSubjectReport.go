package report

import (
	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) GetSubjectReport(
	studentId int64,
	name string,
	course byte,
) ([]float64, *apperrors.Error) {
	averages, err := s.Repo.GetSubjectReport(studentId, name, course)

	if err != nil {
		return averages, apperrors.NewError(err)
	}

	return averages, nil
}
