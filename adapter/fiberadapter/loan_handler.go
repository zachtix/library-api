package fiberadapter

import (
	"library/adapter/httpdto"
	"library/core/domain"
	"library/port/inport"
	"library/port/outport"
	"strings"

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

func (h *FiberLoanHandler) ListByMember(c fiber.Ctx) error {
	memberID, err := parseUUIDParam(c, "id")
	if err != nil {
		return err
	}

	var status *domain.LoanStatus
	if q := c.Query("status"); q != "" {
		s := domain.LoanStatus(strings.ToUpper(q))
		switch s {
		case domain.LoanStatusActive, domain.LoanStatusOverdue, domain.LoanStatusReturned:
			status = &s
		default:
			return &validationError{
				message: "validation failed",
				fields:  map[string]string{"status": "status must be one of active, overdue, returned"},
			}
		}
	}

	loans, err := h.service.ListByMember(memberID, status)
	if err != nil {
		return err
	}

	return newOKResponse(c, httpdto.LoanResponsesFromDomain(loans, h.clock.Now()), "list loans successfully")
}
