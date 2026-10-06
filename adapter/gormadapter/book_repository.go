package gormadapter

import (
	"errors"
	"library/core/domain"
	"library/port/outport"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormBookRepository struct {
	db *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) outport.BookRepository {
	return &GormBookRepository{
		db: db,
	}
}

func (r *GormBookRepository) Save(book domain.Book) (domain.Book, error) {
	model := bookFromDomain(book)
	if err := r.db.Create(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.Book{}, domain.ErrIsbnTaken
		}
		return domain.Book{}, err
	}
	return model.toDomain(), nil
}

func (r *GormBookRepository) FindByID(id uuid.UUID) (domain.Book, error) {
	var model BookModel
	if err := r.db.Take(&model, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Book{}, domain.ErrBookNotFound
		}
		return domain.Book{}, err
	}
	return model.toDomain(), nil
}

func (r *GormBookRepository) AddCopie(copie domain.BookCopie) (domain.BookCopie, error) {
	model := bookCopieFromDomain(copie)
	if err := r.db.Create(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.BookCopie{}, domain.ErrBarcodeTaken
		}
		return domain.BookCopie{}, err
	}
	return model.toDomain(), nil
}

func (r *GormBookRepository) List(page domain.PageRequest, filter domain.BookFilter) ([]domain.Book, int64, error) {
	var models []BookModel

	query := r.db.Model(&BookModel{})
	if q := strings.TrimSpace(filter.Q); q != "" {
		like := "%" + q + "%"
		query = query.Where("(title ILIKE ? OR author ILIKE ?)", like, like)
	}

	query = query.Session(&gorm.Session{})

	if err := query.Order("title").Limit(page.Limit).Offset(page.Offset()).Find(&models).Error; err != nil {
		return []domain.Book{}, 0, err
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return []domain.Book{}, 0, err
	}

	books := make([]domain.Book, 0, len(models))
	for _, m := range models {
		books = append(books, m.toDomain())
	}
	return books, total, nil
}
