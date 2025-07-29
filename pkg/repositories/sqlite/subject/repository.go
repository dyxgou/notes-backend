package subject

import (
	"database/sql"

	"github.com/dyxgou/notas/pkg/ports"
)

var _ ports.SubjectRepository = &Repository{}

type Repository struct {
	Db *sql.DB
}
