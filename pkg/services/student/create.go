package student

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) Create(student *domain.Student) (int64, *domain.AppError) {
	id, err := s.Repo.Insert(student)
	if err != nil {
		return 0, domain.NewError(err)
	}

	return id, nil
}
