package student

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Get(c *fiber.Ctx) error {
	p := new(core.IdParam)

	if err := c.ParamsParser(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	s, err := h.StudentService.Get(p.Id)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(core.StudentResponse{
		Id:          s.Id,
		Name:        s.Name,
		Course:      s.Course,
		ParentPhone: s.ParentPhone,
	})
}
