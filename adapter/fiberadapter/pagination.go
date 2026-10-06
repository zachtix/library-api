package fiberadapter

import (
	"library/core/domain"
	"strings"
)

type PaginationQuery struct {
	Page  int `query:"page" validate:"omitempty,min=1"`
	Limit int `query:"limit" validate:"omitempty,min=1"`
}

func (p *PaginationQuery) toDomain() domain.PageRequest {
	return domain.PageRequest{
		Page:  p.Page,
		Limit: p.Limit,
	}.Normalize()
}

type ListBooksQuery struct {
	PaginationQuery
	Q string `query:"q"`
}

func (q *ListBooksQuery) filterToDomain() domain.BookFilter {
	return domain.BookFilter{Q: strings.TrimSpace(q.Q)}
}
