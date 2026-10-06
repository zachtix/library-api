package fiberadapter

import (
	"library/adapter/httpdto"
	"library/core/domain"

	"github.com/gofiber/fiber/v3"
)

type dataResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func newOKResponse[T any](c fiber.Ctx, data T, message string) error {
	return newDataResponse(c, fiber.StatusOK, data, message)
}

func newCreatedResponse[T any](c fiber.Ctx, data T, message string) error {
	return newDataResponse(c, fiber.StatusCreated, data, message)
}

func newDataResponse[T any](c fiber.Ctx, status int, data T, message string) error {
	return c.Status(status).JSON(dataResponse[T]{
		Data:    data,
		Message: message,
	})
}

func newErrorResponse(c fiber.Ctx, status int, code, message string, fields map[string]string) error {
	return c.Status(status).JSON(httpdto.ErrorResponse{
		Error: httpdto.ErrorBody{Code: code, Message: message, Fields: fields},
	})
}

type paginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

func paginationMetaFromDomain(page domain.PageRequest, total int64) paginationMeta {
	totalPages := 0
	if page.Limit > 0 {
		totalPages = int((total + int64(page.Limit) - 1) / int64(page.Limit))
	}
	return paginationMeta{
		Page:       page.Page,
		Limit:      page.Limit,
		Total:      total,
		TotalPages: totalPages,
	}
}

type paginationResponse[T any] struct {
	Meta    paginationMeta `json:"meta"`
	Data    []T            `json:"data"`
	Message string         `json:"message"`
}

func newPaginationResponse[E, T any](c fiber.Ctx, page domain.Page[E], toDTO func(E) T, message string) error {
	data := make([]T, 0, len(page.Items))
	for _, item := range page.Items {
		data = append(data, toDTO(item))
	}

	return c.Status(fiber.StatusOK).JSON(paginationResponse[T]{
		Meta:    paginationMetaFromDomain(page.Request, page.Total),
		Data:    data,
		Message: message,
	})
}
