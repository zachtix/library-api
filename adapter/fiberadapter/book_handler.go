package fiberadapter

import (
	"library/adapter/httpdto"
	"library/port/inport"

	"github.com/gofiber/fiber/v3"
)

type FiberBookHandler struct {
	service inport.BookService
}

func NewFiberBookHandler(service inport.BookService) *FiberBookHandler {
	return &FiberBookHandler{
		service: service,
	}
}

func (h *FiberBookHandler) Create(c fiber.Ctx) error {
	var body httpdto.CreateBookRequest
	if err := c.Bind().Body(&body); err != nil {
		return newValidationError(err)
	}

	created, err := h.service.Create(body.BookToDomain())
	if err != nil {
		return err
	}
	return newCreatedResponse(c, httpdto.ToBookResponse(created), "book created successfully")
}
func (h *FiberBookHandler) Get(c fiber.Ctx) error {
	id, err := parseUUIDParam(c, "id")
	if err != nil {
		return err
	}

	book, err := h.service.Get(id)
	if err != nil {
		return err
	}

	return newOKResponse(c, httpdto.ToBookResponse(book), "book retrieved successfully")
}
func (h *FiberBookHandler) List(c fiber.Ctx) error {
	var query ListBooksQuery
	if err := c.Bind().Query(&query); err != nil {
		return newValidationError(err)
	}

	page, err := h.service.List(query.toDomain(), query.filterToDomain())
	if err != nil {
		return err
	}
	return newPaginationResponse(c, page, httpdto.ToBookResponse, "books retrieved successfully")
}
func (h *FiberBookHandler) AddCopie(c fiber.Ctx) error {
	id, err := parseUUIDParam(c, "id")
	if err != nil {
		return err
	}

	var body httpdto.CreateBookCopieRequest
	if err := c.Bind().Body(&body); err != nil {
		return newValidationError(err)
	}

	copie, err := h.service.AddCopie(id, body.BookCopieToDomain())
	if err != nil {
		return err
	}
	return newCreatedResponse(c, httpdto.ToBookCopieResponse(copie), "book copie created successfully")
}
