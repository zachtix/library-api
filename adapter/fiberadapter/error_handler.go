package fiberadapter

import (
	"errors"
	"library/core/domain"
	"log"

	"github.com/gofiber/fiber/v3"
)

type validationError struct {
	message string
	fields  map[string]string
}

func (e *validationError) Error() string { return e.message }

func newValidationError(err error) error {
	if fields, ok := validationFields(err); ok {
		return &validationError{message: "validation failed", fields: fields}
	}
	return &validationError{message: "invalid request body"}
}

var domainErrors = []struct {
	err    error
	status int
	code   string
}{
	{domain.ErrMemberNotFound, fiber.StatusNotFound, "MEMBER_NOT_FOUND"},
	{domain.ErrEmailTaken, fiber.StatusConflict, "EMAIL_ALREADY_EXISTS"},
	{domain.ErrInvalidStatus, fiber.StatusBadRequest, "VALIDATION_ERROR"},
	{domain.ErrBookNotFound, fiber.StatusNotFound, "BOOK_NOT_FOUND"},
	{domain.ErrIsbnTaken, fiber.StatusConflict, "ISBN_ALREADY_EXISTS"},
	{domain.ErrBarcodeTaken, fiber.StatusConflict, "BARCODE_ALREADY_EXISTS"},
	{domain.ErrMemberSuspended, fiber.StatusUnprocessableEntity, "MEMBER_SUSPENDED"},
	{domain.ErrHasUnpaidFines, fiber.StatusUnprocessableEntity, "HAS_UNPAID_FINES"},
	{domain.ErrHasOverdueLoans, fiber.StatusUnprocessableEntity, "HAS_OVERDUE_LOANS"},
	{domain.ErrLoanLimitReached, fiber.StatusUnprocessableEntity, "LOAN_LIMIT_REACHED"},
	{domain.ErrCopyNotFound, fiber.StatusNotFound, "COPY_NOT_FOUND"},
	{domain.ErrCopyNotAvailable, fiber.StatusConflict, "COPY_NOT_AVAILABLE"},
	{domain.ErrLoanNotFound, fiber.StatusNotFound, "LOAN_NOT_FOUND"},
	{domain.ErrLoanAlreadyReturned, fiber.StatusConflict, "LOAN_ALREADY_RETURNED"},
	{domain.ErrLoanOverdue, fiber.StatusUnprocessableEntity, "LOAN_OVERDUE"},
	{domain.ErrRenewLimitReached, fiber.StatusUnprocessableEntity, "RENEW_LIMIT_REACHED"},
}

func ErrorHandler(c fiber.Ctx, err error) error {
	var ve *validationError
	if errors.As(err, &ve) {
		return newErrorResponse(c, fiber.StatusBadRequest, "VALIDATION_ERROR", ve.message, ve.fields)
	}

	for _, d := range domainErrors {
		if errors.Is(err, d.err) {
			return newErrorResponse(c, d.status, d.code, err.Error(), nil)
		}
	}

	var fe *fiber.Error
	if errors.As(err, &fe) {
		return newErrorResponse(c, fe.Code, "HTTP_ERROR", fe.Message, nil)
	}

	log.Printf("internal error: %s %s: %v", c.Method(), c.Path(), err)
	return newErrorResponse(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "internal server error", nil)
}
