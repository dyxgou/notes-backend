package grade

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) ChangeName(c *fiber.Ctx) error {
	q := new(core.ChangeGradeName)

	if err := c.BodyParser(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.GradeService.ChangeName(q.Id, q.Name); err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(fiber.Map{"msg": "grade name changed successfully"})
}
