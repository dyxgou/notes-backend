package note

import "github.com/dyxgou/notas/pkg/domain"

func (s *Service) Get(gradeId, studentId int64) (*domain.Note, *domain.AppError) {
	note, err := s.Repo.GetByGradeIdAndStudentId(gradeId, studentId)

	if err != nil {
		return nil, domain.NewError(err)
	}

	return note, nil
}
