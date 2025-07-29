package student

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) ChangeName(id int64, name string) (int64, *domain.AppError) {
	id, err := s.Repo.ChangeName(id, name)

	if err != nil {
		return 0, domain.NewError(err)
	}

	return id, nil
}
