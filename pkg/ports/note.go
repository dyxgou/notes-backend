package ports

import "github.com/dyxgou/notas/pkg/domain"

type NoteService interface {
	CreateAndGet(note *domain.Note) (*domain.Note, *domain.AppError)
	Get(gradeId, studentId int64) (*domain.Note, *domain.AppError)
	ChangeValue(id int64, value byte) *domain.AppError
}

type NoteRepository interface {
	Insert(note *domain.Note) (int64, error)
	GetByGradeIdAndStudentId(gradeId, studentId int64) (*domain.Note, error)
	Get(id int64) (*domain.Note, error)
	ChangeValue(id int64, value byte) error
}
