package ports

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

type SubjectService interface {
	Create(subject *domain.Subject) (int64, *apperrors.Error)
	CreateAndGetSubject(subject *domain.Subject) (int64, *apperrors.Error)
	GetByCourseAndPeriod(course, period byte) ([]domain.Subject, *apperrors.Error)
}

type SubjectRepository interface {
	Insert(subject *domain.Subject) (int64, error)
	GetSubjectById(id int64) (*domain.Subject, error)
	GetByNameAndCourse(name string, course, period byte) (int64, error)
	GetByCourseAndPeriod(course, period byte) ([]domain.Subject, error)
}
