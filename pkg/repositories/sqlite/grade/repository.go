package grade

import (
	"database/sql"

	"github.com/dyxgou/notas/pkg/ports"
)

var _ ports.GradeRespository = &Repository{}

type Repository struct {
	Db *sql.DB
}
