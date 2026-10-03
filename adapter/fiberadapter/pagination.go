package fiberadapter

import "library/core/domain"

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
