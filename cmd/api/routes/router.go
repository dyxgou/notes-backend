package routes

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type Router struct {
	Db       *sql.DB
	Validate *validator.Validate
}

func NewRouter(db *sql.DB, val *validator.Validate) *Router {
	return &Router{Db: db, Validate: val}
}
