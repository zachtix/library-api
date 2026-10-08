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
	if err := r.db.Take(&model, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Loan{}, domain.ErrLoanNotFound
		}
		return domain.Loan{}, err
	}
	return model.toDomain(), nil
}
func (r *GormLoanRepository) Update(loan domain.Loan) (domain.Loan, error) {
	model := r.db.Model(&LoanModel{}).Where("id = ?", loan.ID).Updates(map[string]any{
		"due_at":      loan.DueAt,
		"returned_at": loan.ReturnedAt,
		"renew_count": loan.RenewCount,
	})
	if model.Error != nil {
		return domain.Loan{}, model.Error
	}
	if model.RowsAffected == 0 {
		return domain.Loan{}, domain.ErrLoanNotFound
	}
	return r.FindByID(loan.ID)
}

func (r *GormLoanRepository) ListActiveByMember(memberID uuid.UUID) ([]domain.Loan, error) {
	var models []LoanModel
	if err := r.db.Where("member_id = ? AND returned_at IS NULL", memberID).Order("borrowed_at DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	return loansToDomain(models), nil
}

func (r *GormLoanRepository) ListByMember(memberID uuid.UUID) ([]domain.Loan, error) {
	var models []LoanModel
	if err := r.db.Where("member_id = ?", memberID).Order("borrowed_at DESC").Find(&models).Error; err != nil {
		return nil, err
	}

	return loansToDomain(models), nil
}
