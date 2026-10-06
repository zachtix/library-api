package inport

import (
	"library/core/domain"

	"github.com/google/uuid"
)

type BookService interface {
	Create(book domain.Book) (domain.Book, error)
	List(page domain.PageRequest, filter domain.BookFilter) (domain.Page[domain.Book], error)
	Get(id uuid.UUID) (domain.Book, error)
	AddCopie(id uuid.UUID, copie domain.BookCopie) (domain.BookCopie, error)
}
