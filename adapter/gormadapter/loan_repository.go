package gormadapter

import (
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

func (r *GormLoanRepository) Save(load domain.Loan) (domain.Loan, error)
func (r *GormLoanRepository) FindByID(id uuid.UUID) (domain.Loan, error)
func (r *GormLoanRepository) Update(load domain.Loan) (domain.Loan, error)
func (r *GormLoanRepository) ListActiveByMember(memberID uuid.UUID) ([]domain.Loan, error)
func (r *GormLoanRepository) ListByMember(memberID uuid.UUID) ([]domain.Loan, error)
