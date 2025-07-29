package student

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetParentPhone(c *fiber.Ctx) error {
	s := new(core.IdParam)

	if err := c.ParamsParser(s); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	tel, err := h.StudentService.GetParentPhone(s.Id)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(fiber.Map{
		"parent_phone": tel,
	})
}
