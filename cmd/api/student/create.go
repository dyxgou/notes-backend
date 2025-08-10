package student

import (
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/dyxgou/notas/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) Create(c *fiber.Ctx) error {
	s := new(core.CreateStudentParams)

	if err := c.BodyParser(s); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(s); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	student := &domain.Student{
		Name:        s.Name,
		Course:      s.Course,
		ParentPhone: s.ParentPhone,
	}

	id, err := h.StudentService.Create(student)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	return c.JSON(&core.StudentResponse{
		Id:          id,
		Name:        student.Name,
		ParentPhone: student.ParentPhone,
		Course:      student.Course,
	})
}
