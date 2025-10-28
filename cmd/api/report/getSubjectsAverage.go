package report

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetSubjectsAverage(c *fiber.Ctx) error {
	q := new(core.GetSubjectsAverageParams)

	if err := c.QueryParser(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	average, err := h.ReportService.GetSubjectsAverage(q.StudentId, q.Course, q.Names)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(fiber.Map{
		"average": average,
	})
}
