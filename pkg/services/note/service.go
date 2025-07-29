package note

import "github.com/dyxgou/notas/pkg/ports"

var _ ports.NoteService = &Service{}

type Service struct {
	Repo ports.NoteRepository
}
