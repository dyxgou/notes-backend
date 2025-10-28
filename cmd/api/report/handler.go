package report

import (
	"database/sql"

	"github.com/dyxgou/notas/pkg/ports"
	reportRepo "github.com/dyxgou/notas/pkg/repositories/sqlite/report"
	reportService "github.com/dyxgou/notas/pkg/services/report"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	ReportService ports.ReportService
	Validate      *validator.Validate
}

func NewHandler(db *sql.DB, val *validator.Validate) *Handler {
	return &Handler{
		ReportService: &reportService.Service{
			Repo: &reportRepo.Repository{
				Db: db,
			},
		},

		Validate: val,
	}
}
