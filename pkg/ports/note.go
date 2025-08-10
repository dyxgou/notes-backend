package ports

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

type NoteService interface {
	CreateAndGet(note *domain.Note) (*domain.Note, *apperrors.Error)
	Get(gradeId, studentId int64) (*domain.Note, *apperrors.Error)
	GetAllStudentNotes(studentId, subjectId int64) ([]domain.Note, *apperrors.Error)
	ChangeValue(id int64, value byte) *apperrors.Error
}

type NoteRepository interface {
	Insert(note *domain.Note) (int64, error)
	GetByGradeIdAndStudentId(gradeId, studentId int64) (*domain.Note, error)
	Get(id int64) (*domain.Note, error)
	GetAllStudentNotes(studentId, subjectId int64) ([]domain.Note, error)
	ChangeValue(id int64, value byte) error
}
