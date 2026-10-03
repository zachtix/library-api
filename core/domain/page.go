package domain

const (
	DefaultPageLimit = 10
	MaxPageLimit     = 100
)

type PageRequest struct {
	Page  int
	Limit int
}

func (p PageRequest) Normalize() PageRequest {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = DefaultPageLimit
	}
	if p.Limit > MaxPageLimit {
		p.Limit = MaxPageLimit
	}
	return p
}

func (p PageRequest) Offset() int {
	return (p.Page - 1) * p.Limit
}

type Page[T any] struct {
	Items   []T
	Total   int64
	Request PageRequest
}
