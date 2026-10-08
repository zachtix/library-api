package main

import (
	"library/adapter/clockadapter"
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

	tx := gormadapter.NewGormTx(db)
	clock := clockadapter.NewClock()

	memberRepo := gormadapter.NewGormMemberRepository(db)
	memberService := service.NewMemberService(memberRepo, uuidadapter.NewUUIDGenerator())
	memberHandler := fiberadapter.NewFiberMemberHandler(memberService)
	membersRoute := app.Group("/members")
	membersRoute.Post("", memberHandler.Create)
	membersRoute.Get("/:id", memberHandler.Get)
	membersRoute.Patch("/:id/status", memberHandler.UpdateStatus)

	bookRepo := gormadapter.NewGormBookRepository(db)
	bookService := service.NewBookService(bookRepo, uuidadapter.NewUUIDGenerator())
	bookHandler := fiberadapter.NewFiberBookHandler(bookService)
	booksRoute := app.Group("/books")
	booksRoute.Post("", bookHandler.Create)
	booksRoute.Get("", bookHandler.List)
	booksRoute.Get("/:id", bookHandler.Get)
	booksRoute.Post("/:id/copies", bookHandler.AddCopie)

	loanRepo := gormadapter.NewGormLoanRepository(db)
	loanService := service.NewLoanService(loanRepo, uuidadapter.NewUUIDGenerator(), tx, clock)
	loanHandler := fiberadapter.NewFiberLoanHandler(loanService, clock)
	loanRoute := app.Group("/loans")
	loanRoute.Post("", loanHandler.Borrow)
	loanRoute.Post("/:id/renew", loanHandler.Renew)
	loanRoute.Post("/:id/return", loanHandler.Return)

	log.Fatal(app.Listen(":8080"))
}
