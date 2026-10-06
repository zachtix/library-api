package main

import (
	"library/adapter/fiberadapter"
	"library/adapter/gormadapter"
	"library/adapter/uuidadapter"
	"library/core/service"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		StructValidator: fiberadapter.NewStructValidator(),
		ErrorHandler:    fiberadapter.ErrorHandler,
	})

	app.Get("/healthz", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "ok",
		})
	})

	memberRepo := gormadapter.NewGormMemberRepository(db)
	memberService := service.NewMemberService(memberRepo, uuidadapter.NewUUIDGenerator())
	memberHandler := fiberadapter.NewFiberMemberHandler(memberService)

	membersRoute := app.Group("/members")
	membersRoute.Post("", memberHandler.Create)
	membersRoute.Get("/:id", memberHandler.Get)
	membersRoute.Patch("/:id/status", memberHandler.UpdateStatus)

	log.Fatal(app.Listen(":8080"))
}
