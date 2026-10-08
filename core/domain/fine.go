package domain

import (
	"time"

	"github.com/google/uuid"
)

type Fine struct {
	ID        uuid.UUID
	LoanID    uuid.UUID
	MemberID  uuid.UUID
	Amount    int64
	PaidAt    *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

const (
	FinePerDay int64 = 10
	MaxFine    int64 = 500
)

func CalculateFine(dueAt, returnedAt time.Time) int64 {
	days := LateDays(dueAt, returnedAt)
	if days <= 0 {
		return 0
	}
	return min(int64(days)*FinePerDay, MaxFine)
}
