package core

import "github.com/gofiber/fiber/v2"

func ErrToJSON(err error) *fiber.Map {
	return &fiber.Map{
		"error": err.Error(),
	}
}
