package report

import (
	"database/sql"

	"github.com/dyxgou/notas/pkg/ports"
)

var _ ports.ReportRepository = &Repository{}

type Repository struct {
	Db *sql.DB
}
