package main

import (
	"library/adapter/fiberadapter"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		StructValidator: fiberadapter.NewStructValidator(),
	})

	app.Get("healthz", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "ok",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
