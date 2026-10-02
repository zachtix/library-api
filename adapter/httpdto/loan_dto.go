package httpdto

import (
	"library/core/domain"
	"time"

	"github.com/google/uuid"
)

type CreateLoanRequest struct {
	MemberID string `json:"member_id" validate:"required"`
	Barcode  string `json:"barcode" validate:"required"`
}

type LoanResponse struct {
	ID         uuid.UUID         `json:"id"`
	MemberID   uuid.UUID         `json:"member_id"`
	CopyID     uuid.UUID         `json:"copy_id"`
	BorrowedAt time.Time         `json:"borrowed_at"`
	DueAt      time.Time         `json:"due_at"`
	ReturnedAt *time.Time        `json:"returned_at"`
	RenewCount int               `json:"renew_count"`
	Status     domain.LoanStatus `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

type ReturnLoanResponse struct {
	Loan LoanResponse  `json:"loan"`
	Fine *FineResponse `json:"fine"`
}
