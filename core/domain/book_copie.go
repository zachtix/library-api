package domain

import (
	"time"

	"github.com/google/uuid"
)

type BookStatus string

const (
	CopyStatusAvailable BookStatus = "AVAILABLE"
	CopyStatusBorrowed  BookStatus = "BORROWED"
	CopyStatusLost      BookStatus = "LOST"
)

type BookCopie struct {
	ID        uuid.UUID
	BookID    uuid.UUID
	Barcode   string
	Status    BookStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
