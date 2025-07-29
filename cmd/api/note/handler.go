package note

import (
	"database/sql"

	"github.com/dyxgou/notas/pkg/ports"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/note"
	noteService "github.com/dyxgou/notas/pkg/services/note"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	NoteService ports.NoteService
	Validate    *validator.Validate
}

func NewHandler(db *sql.DB, val *validator.Validate) *Handler {
	return &Handler{
		NoteService: &noteService.Service{
			Repo: &note.Repository{
				Db: db,
			},
		},

		Validate: val,
	}
}
