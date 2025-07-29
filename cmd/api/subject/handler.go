package subject

import (
	"database/sql"

	"github.com/dyxgou/notas/pkg/ports"
	subjectRepo "github.com/dyxgou/notas/pkg/repositories/sqlite/subject"
	subjectService "github.com/dyxgou/notas/pkg/services/subject"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	SubjectService ports.SubjectService
	Validate       *validator.Validate
}

func NewHandler(db *sql.DB, val *validator.Validate) *Handler {
	return &Handler{
		SubjectService: &subjectService.Service{
			Repo: &subjectRepo.Repository{
				Db: db,
			},
		},

		Validate: val,
	}
}
