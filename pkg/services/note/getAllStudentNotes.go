package note

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (s *Service) GetAllStudentNotes(studentId, subjectId int64) ([]domain.Note, *apperrors.Error) {
	n, err := s.Repo.GetAllStudentNotes(studentId, subjectId)

	if err != nil {
		return n, apperrors.NewError(err)
	}

	return n, nil
}
