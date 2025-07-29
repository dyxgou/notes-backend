package grade

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Create(grade *domain.Grade) (int64, *domain.AppError) {
	id, err := s.Repo.Insert(grade)
	if err != nil {
		return 0, domain.NewError(err)
	}

	return id, nil
}
