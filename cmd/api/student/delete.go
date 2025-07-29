package student

import (
	"errors"
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Delete(c *fiber.Ctx) error {
	s := new(core.IdParam)

	if err := c.ParamsParser(s); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	id, err := h.StudentService.Delete(s.Id)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	if id == 0 {
		return c.Status(http.StatusNotFound).JSON(core.ErrToJSON(
			errors.New("invalid student id"),
		))
	}

	return c.JSON(fiber.Map{
		"msg": "user deleted",
	})
}
