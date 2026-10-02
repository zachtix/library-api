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
