package student

import (
	"errors"
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetStudentsByCourse(c *fiber.Ctx) error {
	cs := new(core.IdParam)

	if err := c.ParamsParser(cs); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if cs.Id < 0 || cs.Id > 11 {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(
			errors.New("course should be between 0 and 11"),
		))
	}

	students, err := h.StudentService.GetStudentsByCourse(cs.Id)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(students)
}
