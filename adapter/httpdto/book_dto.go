package httpdto

import (
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
