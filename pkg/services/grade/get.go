package grade

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Get(id int64) (*domain.Grade, *domain.AppError) {
	grade, err := s.Repo.Get(id)

	if err != nil {
		return nil, domain.NewError(err)
	}

	return grade, nil
}
