package student

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetStudentAverage(c *fiber.Ctx) error {
	q := new(core.GetStudentAverageQuery)

	if err := c.QueryParser(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	avg, err := h.StudentService.GetStudentAverage(q.StudentId, q.SubjectId)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(&fiber.Map{"average": avg})
}
