package note

import (
	"database/sql"

	"github.com/dyxgou/notas/pkg/ports"
)

var _ ports.NoteRepository = &Repository{}

type Repository struct {
	Db *sql.DB
}
