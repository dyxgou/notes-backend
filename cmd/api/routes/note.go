package routes

import (
	"github.com/dyxgou/notas/cmd/api/note"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) RegisterNoteGroup(router fiber.Router) {
	h := note.NewHandler(r.Db, r.Validate)

	router.Post("/", h.Create)
	router.Get("/", h.Get)
	router.Patch("/", h.ChangeValue)
}
