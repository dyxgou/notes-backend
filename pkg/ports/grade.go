package ports

import "github.com/dyxgou/notas/pkg/domain"

type GradeService interface {
	Create(grade *domain.Grade) (int64, *domain.AppError)
	Get(id int64) (*domain.Grade, *domain.AppError)
	GetGradesBySubjectId(subjectId int64) ([]domain.Grade, *domain.AppError)
	ChangeName(id int64, name string) *domain.AppError
}

type GradeRespository interface {
	Insert(grade *domain.Grade) (int64, error)
	Get(id int64) (*domain.Grade, error)
	GetGradesBySubjectId(subjectId int64) ([]domain.Grade, error)
	ChangeName(id int64, name string) error
}
