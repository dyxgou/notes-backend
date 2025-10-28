package ports

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

type GradeService interface {
	Create(grade *domain.Grade) (int64, *apperrors.Error)
	Get(id int64) (*domain.Grade, *apperrors.Error)
	Delete(gradeId, subjectId int64) (int64, *apperrors.Error)
	GetGradesBySubjectId(subjectId int64) ([]domain.Grade, *apperrors.Error)
	ChangeName(id int64, name string) *apperrors.Error
}

type GradeRespository interface {
	Insert(grade *domain.Grade) (int64, error)
	Get(id int64) (*domain.Grade, error)
	Delete(gradeId, subjectId int64) (int64, error)
	GetGradesBySubjectId(subjectId int64) ([]domain.Grade, error)
	ChangeName(id int64, name string) error
}
