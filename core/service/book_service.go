package service

import (
	"library/core/domain"
	"library/port/inport"
	"library/port/outport"

	"github.com/google/uuid"
)

type bookServiceImpl struct {
	repo  outport.BookRepository
	idGen outport.IDGenerator
}

func NewBookService(repo outport.BookRepository, idGen outport.IDGenerator) inport.BookService {
	return &bookServiceImpl{
		repo:  repo,
		idGen: idGen,
	}
}

func (s *bookServiceImpl) Create(book domain.Book) (domain.Book, error) {
	book.ID = s.idGen.NewID()

	created, err := s.repo.Save(book)
	if err != nil {
		return domain.Book{}, err
	}
	return created, nil
}
func (s *bookServiceImpl) Get(id uuid.UUID) (domain.Book, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}
func (s *bookServiceImpl) List(page domain.PageRequest, filter domain.BookFilter) (domain.Page[domain.Book], error) {
	page = page.Normalize()
	books, total, err := s.repo.List(page, filter)
	if err != nil {
		return domain.Page[domain.Book]{}, err
	}
	return domain.Page[domain.Book]{Items: books, Total: total, Request: page}, nil
}
func (s *bookServiceImpl) AddCopie(id uuid.UUID, copie domain.BookCopie) (domain.BookCopie, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return domain.BookCopie{}, err
	}

	copie.ID = s.idGen.NewID()
	copie.BookID = book.ID
	copie.Status = domain.CopyStatusAvailable

	created, err := s.repo.AddCopie(copie)
	if err != nil {
		return domain.BookCopie{}, err
	}
	return created, nil
}
