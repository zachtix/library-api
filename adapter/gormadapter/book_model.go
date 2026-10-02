package gormadapter

import (
	"library/core/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;not null"`
	Isbn      string    `gorm:"uniqueIndex:idx_books_isbn,where:deleted_at IS NULL;not null"`
	Title     string    `gorm:"not null"`
	Author    string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (BookModel) TableName() string { return "books" }

func (m BookModel) toDomain() domain.Book {
	return domain.Book{
		ID:        m.ID,
		Isbn:      m.Isbn,
		Title:     m.Title,
		Author:    m.Author,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: deletedAtToDomain(m.DeletedAt),
	}
}

func bookFromDomain(b domain.Book) BookModel {
	return BookModel{
		ID:        b.ID,
		Isbn:      b.Isbn,
		Title:     b.Title,
		Author:    b.Author,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		DeletedAt: deletedAtFromDomain(b.DeletedAt),
	}
}
