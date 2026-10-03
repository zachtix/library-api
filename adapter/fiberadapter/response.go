package fiberadapter

import (
	"library/core/domain"

	"github.com/gofiber/fiber/v3"
)

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
