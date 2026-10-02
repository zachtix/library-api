package domain

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID
	Isbn      string
	Title     string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
