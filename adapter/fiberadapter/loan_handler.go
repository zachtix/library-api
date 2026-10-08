package fiberadapter

import (
	"library/adapter/httpdto"
	"library/port/inport"
	"library/port/outport"

	"github.com/gofiber/fiber/v3"
)

type FiberLoanHandler struct {
	service inport.LoanService
	clock   outport.Clock
}

func NewFiberLoanHandler(service inport.LoanService, clock outport.Clock) *FiberLoanHandler {
	return &FiberLoanHandler{
		service: service,
		clock:   clock,
	}
}

func (h *FiberLoanHandler) Borrow(c fiber.Ctx) error {
	var body httpdto.CreateLoanRequest
	if err := c.Bind().Body(&body); err != nil {
		return newValidationError(err)
	}

	created, err := h.service.Borrow(body.MemberID, body.Barcode)
	if err != nil {
		return err
	}

	return newCreatedResponse(c, httpdto.LoanResponseFromDomain(created, h.clock.Now()), "borrow created successfully")
}

func (h *FiberLoanHandler) Renew(c fiber.Ctx) error {
	id, err := parseUUIDParam(c, "id")
	if err != nil {
		return err
	}

	renew, err := h.service.Renew(id)
	if err != nil {
		return err
	}

	return newOKResponse(c, httpdto.LoanResponseFromDomain(renew, h.clock.Now()), "borrow renewed successfully")
}

func (h *FiberLoanHandler) Return(c fiber.Ctx) error {
	id, err := parseUUIDParam(c, "id")
	if err != nil {
		return err
	}

	ret, fine, err := h.service.Return(id)
	if err != nil {
		return err
	}

	return newDataResponse(c, fiber.StatusOK, httpdto.ReturnLoanResponseFromDomain(ret, fine, h.clock.Now()), "return borrow successfully")
}
