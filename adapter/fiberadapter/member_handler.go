package fiberadapter

import (
	"errors"
	"library/adapter/httpdto"
	"library/core/domain"
	"library/port/inport"

	"github.com/gofiber/fiber/v3"
)

type FiberMemberHandler struct {
	service inport.MemberService
}

func NewFiberMemberHandler(service inport.MemberService) *FiberMemberHandler {
	return &FiberMemberHandler{
		service: service,
	}
}

func (h *FiberMemberHandler) CreateMember(c fiber.Ctx) error {
	var body httpdto.CreateMemberRequest
	if err := c.Bind().Body(&body); err != nil {
		return newErrorResponse(c, fiber.StatusBadRequest, err)
	}

	created, err := h.service.Create(body.MemberToDomain())
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrEmailTaken):
			return newErrorResponse(c, fiber.StatusConflict, err)
		default:
			return newErrorResponse(c, fiber.StatusInternalServerError, err)
		}
	}

	return newOKResponse(c, httpdto.MemberResponseFromDomain(created), "member created successfully")
}
