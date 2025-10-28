package report

import (
	"math"

	"github.com/dyxgou/notas/pkg/apperrors"
)

func (s *Service) GetSubjectsAverage(
	studentId int64,
	course byte,
	names []string,
) (float64, *apperrors.Error) {
	var average float64

	for _, name := range names {
		avg, err := s.Repo.GetSubjectAverage(studentId, name, course)
		if err != nil {
			return average, apperrors.NewError(err)
		}

		average += avg
	}

	if len(names) == 0 {
		return 10.0, nil
	}

	rnd := math.Round(average / float64(len(names)))
	return rnd, nil
}
