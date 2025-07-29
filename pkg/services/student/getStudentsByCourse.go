package student

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) GetStudentsByCourse(courseId int64) ([]domain.Student, *domain.AppError) {
	students, err := s.Repo.GetStudentsByCourse(courseId)

	if err != nil {
		return students, domain.NewError(err)
	}

	return students, nil
}
