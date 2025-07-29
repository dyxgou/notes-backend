package ports

import "github.com/dyxgou/notas/pkg/domain"

type StudentService interface {
	Create(student *domain.Student) (int64, *domain.AppError)
	Get(id int64) (*domain.Student, *domain.AppError)
	GetParentPhone(id int64) (string, *domain.AppError)
	ChangeName(id int64, name string) (int64, *domain.AppError)
	ChangeParentPhone(id int64, phone string) (int64, *domain.AppError)
	GetStudentsByCourse(courseId int64) ([]domain.Student, *domain.AppError)
	Delete(id int64) (int64, *domain.AppError)
}

type StudentRepository interface {
	Insert(student *domain.Student) (int64, error)
	Get(id int64) (*domain.Student, error)
	GetParentPhone(id int64) (string, error)
	GetStudentsByCourse(courseId int64) ([]domain.Student, error)
	ChangeName(id int64, name string) (int64, error)
	ChangeParentPhone(id int64, phone string) (int64, error)
	Delete(id int64) (int64, error)
}
