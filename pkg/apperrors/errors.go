package apperrors

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-sqlite3"
)

type Error struct {
	Msg    error
	Status int
}

func NewError(err error) *Error {
	e := &Error{Msg: err}
	e.setStatus()

	return e
}

func (e *Error) setStatus() {
	e.Status = http.StatusBadRequest

	var sqlErr sqlite3.Error
	if errors.As(e.Msg, &sqlErr) {
		if sqlErr.Code == sqlite3.ErrConstraint {
			e.Status = http.StatusConflict
		}

		slog.Info("sqlite error", "err", sqlite3.ErrConstraint)
		slog.Info("sqlerr", "expanded", sqlErr.Code)
	}

	if errors.Is(e.Msg, sql.ErrNoRows) {
		e.Status = http.StatusNotFound
	}

	if errors.Is(e.Msg, ErrSubjectHasFinalExam) {
		e.Status = http.StatusConflict
	}
}

func (e *Error) Error() string {
	return e.Msg.Error()
}

func (e *Error) ToJSON() *fiber.Map {
	return &fiber.Map{
		"error": e.Msg.Error(),
	}
}
