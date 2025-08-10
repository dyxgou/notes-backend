package grade

import (
	"github.com/dyxgou/notas/pkg/apperrors"
	"github.com/dyxgou/notas/pkg/domain"
)

func (r *Service) GetGradesBySubjectId(
	subjectId int64,
) ([]domain.Grade, *apperrors.Error) {
	grades, err := r.Repo.GetGradesBySubjectId(subjectId)

	if err != nil {
		return grades, apperrors.NewError(err)
	}

	return grades, nil
}
