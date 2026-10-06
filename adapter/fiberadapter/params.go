package fiberadapter

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func parseUUIDParam(c fiber.Ctx, name string) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Params(name))
	if err != nil {
		return uuid.Nil, &validationError{
			message: "validation failed",
			fields:  map[string]string{name: name + " must be a valid UUID"},
		}
	}
	return id, nil
}
