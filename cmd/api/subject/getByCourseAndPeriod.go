package subject

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetByCourseAndPeriod(c *fiber.Ctx) error {
	q := new(core.GetByPeriodAndCourseQuery)

	if err := c.QueryParser(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	s, err := h.SubjectService.GetByCourseAndPeriod(q.Course, q.Period)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(s)
}
