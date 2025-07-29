package subject

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Create(subject *domain.Subject) (int64, *domain.AppError) {
	id, err := s.Repo.Insert(subject)

	if err != nil {
		return 0, domain.NewError(err)
	}

	return id, nil
}
