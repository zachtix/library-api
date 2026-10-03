package httpdto

import (
	"library/core/domain"
	"time"

	"github.com/google/uuid"
)

type CreateBookRequest struct {
	Isbn   string `json:"isbn" validate:"required,numeric,len=13"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
}

type BookResponse struct {
	ID        uuid.UUID `json:"id"`
	Isbn      string    `json:"isbn"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToBookResponse(b domain.Book) BookResponse {
	return BookResponse{
		ID:        b.ID,
		Isbn:      b.Isbn,
		Title:     b.Title,
		Author:    b.Author,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}
