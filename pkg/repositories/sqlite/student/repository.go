package student

import (
	"database/sql"

	"github.com/dyxgou/notas/pkg/ports"
)

var _ ports.StudentRepository = &Repository{}

type Repository struct {
	Db *sql.DB
}
