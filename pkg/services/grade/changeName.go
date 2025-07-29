package grade

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) ChangeName(id int64, name string) *domain.AppError {
	err := s.Repo.ChangeName(id, name)

	if err != nil {
		return domain.NewError(err)
	}

	return nil
}
