package grade

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/dyxgou/notas/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	p := new(core.CreateGradeParams)

	if err := c.BodyParser(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	grade := &domain.Grade{
		Name:        p.Name,
		SubjectId:   p.SubjectId,
		IsFinalExam: p.IsFinalExam,
	}

	id, err := h.GradeService.Create(grade)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}
