package grade

import "github.com/dyxgou/notas/pkg/domain"

func (r *Service) GetGradesBySubjectId(
	subjectId int64,
) ([]domain.Grade, *domain.AppError) {
	grades, err := r.Repo.GetGradesBySubjectId(subjectId)

	if err != nil {
		return grades, domain.NewError(err)
	}

	return grades, nil
}
