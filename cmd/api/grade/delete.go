package grade

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/dyxgou/notas/cmd/api/core"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Delete(c *fiber.Ctx) error {
	p := new(core.DeleteGradeParams)

	if err := c.QueryParser(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	if err := h.Validate.Struct(p); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(core.ErrToJSON(err))
	}

	id, err := h.GradeService.Delete(p.GradeId, p.SubjectId)
	if err != nil {
		return c.Status(err.Status).JSON(err.ToJSON())
	}

	slog.Info("deleting grade", "id", id)

	if id == 0 {
		return c.Status(http.StatusNotFound).JSON(core.ErrToJSON(
			errors.New("invalid grade id"),
		))
	}

	return c.JSON(fiber.Map{
		"msg": "grade was deleted successfully",
	})
}
