package domain

import (
	"errors"
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

var (
	ErrBookNotFound = errors.New("book not found")
	ErrIsbnTaken    = errors.New("isbn already taken")
)

type BookFilter struct {
	Q string
}
