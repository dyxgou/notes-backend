package student

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) GetStudentsByCourse(courseId int64) ([]domain.Student, *apperrors.Error) {
	students, err := s.Repo.GetStudentsByCourse(courseId)

	if err != nil {
		return students, apperrors.NewError(err)
	}

	return students, nil
}
