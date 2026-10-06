package httpdto

import (
	"library/core/domain"
	"time"

	"github.com/google/uuid"
)

type CreateBookCopieRequest struct {
	Barcode string `json:"barcode" validate:"required"`
}

type BookCopieResponse struct {
	ID        uuid.UUID         `json:"id"`
	BookID    uuid.UUID         `json:"book_id"`
	Barcode   string            `json:"barcode"`
	Status    domain.BookStatus `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func (r CreateBookCopieRequest) BookCopieToDomain() domain.BookCopie {
	return domain.BookCopie{
		Barcode: r.Barcode,
	}
}

func ToBookCopieResponse(bc domain.BookCopie) BookCopieResponse {
	return BookCopieResponse{
		ID:        bc.ID,
		BookID:    bc.BookID,
		Barcode:   bc.Barcode,
		Status:    bc.Status,
		CreatedAt: bc.CreatedAt,
		UpdatedAt: bc.UpdatedAt,
	}
}
