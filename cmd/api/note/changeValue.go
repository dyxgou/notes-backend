package note

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) ChangeValue(c *fiber.Ctx) error {
	p := new(core.ChangeNoteValue)

	if err := c.BodyParser(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.NoteService.ChangeValue(p.Id, p.Value); err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(fiber.Map{
		"msg": "note value changed succesfully",
	})
}
