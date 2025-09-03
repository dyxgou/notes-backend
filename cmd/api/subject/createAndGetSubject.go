package subject

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/dyxgou/notas/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateAndGetSubject(c *fiber.Ctx) error {
	s := new(core.CreateSubjectParams)

	if err := c.BodyParser(s); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(s); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	subject := &domain.Subject{
		Name:   s.Name,
		Course: s.Course,
		Period: s.Period,
	}

	id, err := h.SubjectService.CreateAndGetSubject(subject)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(core.IdBody{Id: id})
}
