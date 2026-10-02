package gormadapter

import (
	"library/core/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FineModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;not null"`
	LoanID    uuid.UUID `gorm:"type:uuid;index;not null;uniqueIndex:idx_fines_loan_id,where:deleted_at IS NULL"`
	MemberID  uuid.UUID `gorm:"type:uuid;not null;index"`
	Amount    int64     `gorm:"not null"`
	PaidAt    *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (FineModel) TableName() string { return "fines" }

func (m FineModel) toDomain() domain.Fine {
	return domain.Fine{
		ID:        m.ID,
		LoanID:    m.LoanID,
		MemberID:  m.MemberID,
		Amount:    m.Amount,
		PaidAt:    m.PaidAt,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: deletedAtToDomain(m.DeletedAt),
	}
}

func fineFromDomain(f domain.Fine) FineModel {
	return FineModel{
		ID:        f.ID,
		LoanID:    f.LoanID,
		MemberID:  f.MemberID,
		Amount:    f.Amount,
		PaidAt:    f.PaidAt,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		DeletedAt: deletedAtFromDomain(f.DeletedAt),
	}
}
