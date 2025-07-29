package student

import (
	"database/sql"

	"github.com/dyxgou/notas/pkg/ports"
	studentRepo "github.com/dyxgou/notas/pkg/repositories/sqlite/student"
	studentService "github.com/dyxgou/notas/pkg/services/student"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	StudentService ports.StudentService
	Validate       *validator.Validate
}

func NewHandler(db *sql.DB, val *validator.Validate) *Handler {
	return &Handler{
		StudentService: &studentService.Service{
			Repo: &studentRepo.Repository{
				Db: db,
			},
		},

		Validate: val,
	}
}
