package outport

import (
	"library/core/domain"

	"github.com/google/uuid"
)

type BookRepository interface {
	Save(book domain.Book) (domain.Book, error)
	FindByID(id uuid.UUID) (domain.Book, error)
	List(page domain.PageRequest, filter domain.BookFilter) ([]domain.Book, int64, error)
	AddCopie(copie domain.BookCopie) (domain.BookCopie, error)
	FindCopieByBarcode(barcode string) (domain.BookCopie, error)
	UpdateCopieStatus(id uuid.UUID, status domain.BookStatus) error
}
