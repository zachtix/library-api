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

// var q PaginationQuery
// if err := c.Bind().Query(&q); err != nil {
// 	return err
// }
// page, err := h.svc.ListBooks(c.Context(), q.toDomain())
// if err != nil {
// 	return err
// }

//----------------------------------------------------------
// return newPaginationResponse(c, page, httpdto.ToBookResponse, "ok")

//----------------------------------------------------------
// data := make([]httpdto.BookResponse, 0, len(page.Items))
// for _, b := range page.Items {
// 	data = append(data, httpdto.ToBookResponse(b))
// }

// meta := paginationMetaFromDomain(page.Request, page.Total)
// return newPaginationResponse(c, meta, data, "ok")
