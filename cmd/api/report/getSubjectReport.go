package report

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetSubjectReport(c *fiber.Ctx) error {
	q := new(core.GetSubjectsReportParams)

	if err := c.QueryParser(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(q); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	averages, err := h.ReportService.GetSubjectReport(q.StudentId, q.Name, q.Course)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	if len(averages) == 0 {
		return c.JSON([4]int8{})
	}

	return c.JSON(averages)
}
