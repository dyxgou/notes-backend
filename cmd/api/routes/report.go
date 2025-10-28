package routes

import (
	"github.com/dyxgou/notas/cmd/api/report"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) RegisterReportGroup(router fiber.Router) {
	h := report.NewHandler(r.Db, r.Validate)

	router.Get("/get/", h.GetSubjectReport)
	router.Get("/avg/", h.GetSubjectsAverage)
}
