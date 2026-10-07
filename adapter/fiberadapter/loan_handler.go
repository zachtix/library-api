package fiberadapter

import (
	"library/port/inport"

	"github.com/gofiber/fiber/v3"
)

type FiberLoanHandler struct {
	service inport.LoanService
}

func NewFiberLoanHandler(service inport.LoanService) *FiberLoanHandler {
	return &FiberLoanHandler{
		service: service,
	}
}

func (h *FiberLoanHandler) Borrow(c fiber.Ctx) error

func (h *FiberLoanHandler) Renew(c fiber.Ctx) error

func (h *FiberLoanHandler) Return(c fiber.Ctx) error
