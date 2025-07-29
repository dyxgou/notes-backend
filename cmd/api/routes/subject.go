package routes

import (
	"github.com/dyxgou/notas/cmd/api/subject"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) RegisterSubjectGroup(router fiber.Router) {
	h := subject.NewHandler(r.Db, r.Validate)

	router.Post("/", h.Create)
	router.Post("/get/", h.CreateAndGetSubject)
	router.Get("/", h.GetByNameAndCourse)
}
