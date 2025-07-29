package student

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) GetParentPhone(id int64) (string, *domain.AppError) {
	tel, err := s.Repo.GetParentPhone(id)
	if err != nil {
		return tel, domain.NewError(err)
	}

	return tel, nil
}
