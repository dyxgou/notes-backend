package subject

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) GetByNameAndCourse(name string, course, period byte) (*domain.Subject, *domain.AppError) {
	subject, err := s.Repo.GetByNameAndCourse(name, course, period)

	if err != nil {
		return nil, domain.NewError(err)
	}

	return subject, nil
}
