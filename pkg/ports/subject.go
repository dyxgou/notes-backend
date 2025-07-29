package ports

import "github.com/dyxgou/notas/pkg/domain"

type SubjectService interface {
	Create(subject *domain.Subject) (int64, *domain.AppError)
	CreateAndGetSubject(subject *domain.Subject) (*domain.Subject, *domain.AppError)
	GetByNameAndCourse(name string, course, period byte) (*domain.Subject, *domain.AppError)
}

type SubjectRepository interface {
	Insert(subject *domain.Subject) (int64, error)
	GetSubjectById(id int64) (*domain.Subject, error)
	GetByNameAndCourse(name string, course, period byte) (*domain.Subject, error)
}
