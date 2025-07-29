package student

import (
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) ChangeParentPhone(id int64, phone string) (int64, *domain.AppError) {
	id, err := s.Repo.ChangeParentPhone(id, phone)

	if err != nil {
		return 0, domain.NewError(err)
	}

	return id, nil
}
