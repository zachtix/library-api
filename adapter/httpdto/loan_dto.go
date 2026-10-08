package httpdto

import (
	"library/core/domain"
	"time"

	"github.com/google/uuid"
)

type CreateLoanRequest struct {
	MemberID uuid.UUID `json:"member_id" validate:"required"`
	Barcode  string    `json:"barcode" validate:"required"`
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

func LoanResponseFromDomain(l domain.Loan, now time.Time) LoanResponse {
	return LoanResponse{
		ID:         l.ID,
		MemberID:   l.MemberID,
		CopyID:     l.CopyID,
		BorrowedAt: l.BorrowedAt,
		DueAt:      l.DueAt,
		ReturnedAt: l.ReturnedAt,
		RenewCount: l.RenewCount,
		Status:     l.Status(now),
		CreatedAt:  l.CreatedAt,
		UpdatedAt:  l.UpdatedAt,
	}
}

func LoanResponsesFromDomain(loans []domain.Loan, now time.Time) []LoanResponse {
	res := make([]LoanResponse, 0, len(loans))
	for _, l := range loans {
		res = append(res, LoanResponseFromDomain(l, now))
	}
	return res
}

type ReturnLoanResponse struct {
	Loan LoanResponse  `json:"loan"`
	Fine *FineResponse `json:"fine"`
}

func ReturnLoanResponseFromDomain(l domain.Loan, f *domain.Fine, now time.Time) ReturnLoanResponse {
	res := ReturnLoanResponse{Loan: LoanResponseFromDomain(l, now)}
	if f != nil {
		fine := FineResponseFromDomain(*f)
		res.Fine = &fine
	}
	return res
}
