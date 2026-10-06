package fiberadapter

import (
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

func (h *FiberMemberHandler) Create(c fiber.Ctx) error {
	var body httpdto.CreateMemberRequest
	if err := c.Bind().Body(&body); err != nil {
		return newValidationError(err)
	}

	created, err := h.service.Create(body.MemberToDomain())
	if err != nil {
		return err
	}

	return newCreatedResponse(c, httpdto.MemberResponseFromDomain(created), "member created successfully")
}

func (h *FiberMemberHandler) Get(c fiber.Ctx) error {
	id, err := parseUUIDParam(c, "id")
	if err != nil {
		return err
	}

	member, err := h.service.Get(id)
	if err != nil {
		return err
	}

	return newOKResponse(c, httpdto.MemberResponseFromDomain(member), "member retrieved successfully")
}

func (h *FiberMemberHandler) UpdateStatus(c fiber.Ctx) error {
	id, err := parseUUIDParam(c, "id")
	if err != nil {
		return err
	}

	var body httpdto.UpdateMemberStatusRequest
	if err := c.Bind().Body(&body); err != nil {
		return newValidationError(err)
	}

	updated, err := h.service.UpdateStatus(id, domain.MemberStatus(body.Status))
	if err != nil {
		return err
	}

	return newOKResponse(c, httpdto.MemberResponseFromDomain(updated), "member status updated successfully")
}
