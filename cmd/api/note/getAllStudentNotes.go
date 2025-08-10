package note

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetAllStudentNotes(c *fiber.Ctx) error {
	p := new(core.GetAllStudentNotes)

	if err := c.QueryParser(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	notes, err := h.NoteService.GetAllStudentNotes(p.StudentId, p.SubjectId)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(notes)
}
