package httpdto

import (
	"time"

	"github.com/google/uuid"
)

type FineResponse struct {
	ID        uuid.UUID  `json:"id"`
	LoanID    uuid.UUID  `json:"loan_id"`
	MemberID  uuid.UUID  `json:"member_id"`
	Amount    int64      `json:"amount"`
	PaidAt    *time.Time `json:"paid_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
