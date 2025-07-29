package subject

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetByNameAndCourse(c *fiber.Ctx) error {
	q := new(core.GetByNameAndCourseQuery)

	if err := c.QueryParser(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	subject, err := h.SubjectService.GetByNameAndCourse(q.Name, q.Course, q.Period)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(subject)
}
