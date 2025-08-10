package ports

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

type StudentService interface {
	Create(student *domain.Student) (int64, *apperrors.Error)
	Get(id int64) (*domain.Student, *apperrors.Error)
	GetParentPhone(id int64) (string, *apperrors.Error)
	GetStudentAverage(studentId, subjectId int64) (float64, *apperrors.Error)
	ChangeName(id int64, name string) (int64, *apperrors.Error)
	ChangeParentPhone(id int64, phone string) (int64, *apperrors.Error)
	GetStudentsByCourse(courseId int64) ([]domain.Student, *apperrors.Error)
	Delete(id int64) (int64, *apperrors.Error)
}

type StudentRepository interface {
	Insert(student *domain.Student) (int64, error)
	Get(id int64) (*domain.Student, error)
	GetParentPhone(id int64) (string, error)
	GetStudentsByCourse(courseId int64) ([]domain.Student, error)
	GetStudentAverage(studentId, subjectId int64) (float64, error)
	ChangeName(id int64, name string) (int64, error)
	ChangeParentPhone(id int64, phone string) (int64, error)
	Delete(id int64) (int64, error)
}
