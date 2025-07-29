package note

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/dyxgou/notas/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

var defaultValue byte = 10

func (h *Handler) Create(c *fiber.Ctx) error {
	p := new(core.CreateNoteParams)

	if err := c.BodyParser(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	note := &domain.Note{
		GradeId:   p.GradeId,
		StudentId: p.StudentId,
		Value:     defaultValue,
	}

	n, err := h.NoteService.CreateAndGet(note)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(n)
}
