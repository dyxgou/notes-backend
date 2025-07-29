package routes

import (
	"github.com/dyxgou/notas/cmd/api/student"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) RegisterUserGroup(router fiber.Router) {
	h := student.NewHandler(r.Db, r.Validate)

	router.Post("/", h.Create)
	router.Delete("/:id", h.Delete)
	router.Patch("/change/name", h.ChangeName)
	router.Patch("/change/phone", h.ChangeParentPhone)
	router.Get("/:id", h.Get)
	router.Get("/course/:id", h.GetStudentsByCourse)
	router.Get("/parent/:id", h.GetParentPhone)
}
