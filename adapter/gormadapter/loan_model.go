package gormadapter

import (
	"library/core/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoanModel struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;not null"`
	MemberID   uuid.UUID `gorm:"type:uuid;index;not null"`
	CopyID     uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_loans_active_copy,where:returned_at IS NULL"`
	BorrowedAt time.Time `gorm:"not null"`
	DueAt      time.Time `gorm:"not null"`
	ReturnedAt *time.Time
	RenewCount int `gorm:"default:0;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (LoanModel) TableName() string { return "loans" }

func (m LoanModel) toDomain() domain.Loan {
	return domain.Loan{
		ID:         m.ID,
		MemberID:   m.MemberID,
		CopyID:     m.CopyID,
		BorrowedAt: m.BorrowedAt,
		DueAt:      m.DueAt,
		ReturnedAt: m.ReturnedAt,
		RenewCount: m.RenewCount,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  deletedAtToDomain(m.DeletedAt),
	}
}

func loanFromDomain(l domain.Loan) LoanModel {
	return LoanModel{
		ID:         l.ID,
		MemberID:   l.MemberID,
		CopyID:     l.CopyID,
		BorrowedAt: l.BorrowedAt,
		DueAt:      l.DueAt,
		ReturnedAt: l.ReturnedAt,
		RenewCount: l.RenewCount,
		CreatedAt:  l.CreatedAt,
		UpdatedAt:  l.UpdatedAt,
		DeletedAt:  deletedAtFromDomain(l.DeletedAt),
	}
}

func loansToDomain(models []LoanModel) []domain.Loan {
	loans := make([]domain.Loan, 0, len(models))
	for _, m := range models {
		loans = append(loans, m.toDomain())
	}
	return loans
}
