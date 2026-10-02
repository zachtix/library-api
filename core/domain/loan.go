package domain

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID         uuid.UUID
	MemberID   uuid.UUID
	CopyID     uuid.UUID
	BorrowedAt time.Time
	DueAt      time.Time
	ReturnedAt *time.Time
	RenewCount int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type LoanStatus string

const (
	LoanStatusActive   LoanStatus = "ACTIVE"
	LoanStatusOverdue  LoanStatus = "OVERDUE"
	LoanStatusReturned LoanStatus = "RETURNED"
)

func (l Loan) Status(now time.Time) LoanStatus {
	if l.ReturnedAt != nil {
		return LoanStatusReturned
	}
	if now.After(l.DueAt) {
		return LoanStatusOverdue
	}
	return LoanStatusActive
}
