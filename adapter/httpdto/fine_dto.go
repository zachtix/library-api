package httpdto

import (
	"library/core/domain"
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

func FineResponseFromDomain(f domain.Fine) FineResponse {
	return FineResponse{
		ID:        f.ID,
		LoanID:    f.LoanID,
		MemberID:  f.MemberID,
		Amount:    f.Amount,
		PaidAt:    f.PaidAt,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
	}
}
