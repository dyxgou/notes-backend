package note

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) Get(gradeId, studentId int64) (*domain.Note, *apperrors.Error) {
	note, err := s.Repo.GetByGradeIdAndStudentId(gradeId, studentId)

	if err != nil {
		return nil, apperrors.NewError(err)
	}

	return note, nil
}
