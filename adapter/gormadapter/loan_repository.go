package gormadapter

import (
	"errors"
	"library/core/domain"
	"library/port/outport"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormLoanRepository struct {
	db *gorm.DB
}

func NewGormLoanRepository(db *gorm.DB) outport.LoanRepository {
	return &GormLoanRepository{
		db: db,
	}
}

func (r *GormLoanRepository) Save(loan domain.Loan) (domain.Loan, error) {
	model := loanFromDomain(loan)
	if err := r.db.Create(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.Loan{}, domain.ErrCopyNotAvailable
		}
		return domain.Loan{}, err
	}

	return model.toDomain(), nil
}
func (r *GormLoanRepository) FindByID(id uuid.UUID) (domain.Loan, error) {
	var model LoanModel
	if err := r.db.Where(&model, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Loan{}, domain.ErrLoanNotFound
		}
		return domain.Loan{}, err
	}
	return model.toDomain(), nil
}
func (r *GormLoanRepository) Update(loan domain.Loan) (domain.Loan, error)
func (r *GormLoanRepository) ListActiveByMember(memberID uuid.UUID) ([]domain.Loan, error)
func (r *GormLoanRepository) ListByMember(memberID uuid.UUID) ([]domain.Loan, error)
