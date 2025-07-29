package grade

import (
	"database/sql"

	"github.com/dyxgou/notas/pkg/ports"
	gradeRepo "github.com/dyxgou/notas/pkg/repositories/sqlite/grade"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/subject"
	gradeService "github.com/dyxgou/notas/pkg/services/grade"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	GradeService ports.GradeService
	SubjectRepo  ports.SubjectRepository
	Validate     *validator.Validate
}

func NewHandler(db *sql.DB, val *validator.Validate) *Handler {
	return &Handler{
		GradeService: &gradeService.Service{
			Repo: &gradeRepo.Repository{
				Db: db,
			},
		},

		SubjectRepo: &subject.Repository{
			Db: db,
		},

		Validate: val,
	}
}
