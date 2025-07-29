package student

import (
	"fmt"
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) ChangeName(c *fiber.Ctx) error {
	s := new(core.ChangeStudentName)

	if err := c.BodyParser(s); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(s); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	id, err := h.StudentService.ChangeName(s.Id, s.Name)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	if id == 0 {
		return c.Status(404).JSON(core.ErrToJSON(
			fmt.Errorf("student does not exists"),
		))
	}

	return c.JSON(fiber.Map{
		"msg": "student name changed successfully",
	})
}
