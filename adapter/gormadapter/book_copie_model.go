package gormadapter

import (
	"library/core/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookCopieModel struct {
	ID        uuid.UUID         `gorm:"primaryKey;type:uuid;not null"`
	BookID    uuid.UUID         `gorm:"type:uuid;not null;index"`
	Barcode   string            `gorm:"uniqueIndex:idx_book_copies_barcode,where:deleted_at IS NULL;not null"`
	Status    domain.BookStatus `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (BookCopieModel) TableName() string { return "book_copies" }

func (m BookCopieModel) toDomain() domain.BookCopie {
	return domain.BookCopie{
		ID:        m.ID,
		BookID:    m.BookID,
		Barcode:   m.Barcode,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: deletedAtToDomain(m.DeletedAt),
	}
}

func bookCopieFromDomain(bc domain.BookCopie) BookCopieModel {
	return BookCopieModel{
		ID:        bc.ID,
		BookID:    bc.BookID,
		Barcode:   bc.Barcode,
		Status:    bc.Status,
		CreatedAt: bc.CreatedAt,
		UpdatedAt: bc.UpdatedAt,
		DeletedAt: deletedAtFromDomain(bc.DeletedAt),
	}
}
