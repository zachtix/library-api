package inport

import (
	"library/core/domain"

	"github.com/google/uuid"
)

type LoanService interface {
	Borrow(memberID uuid.UUID, barcode string) (domain.Loan, error)
	Renew(id uuid.UUID) (domain.Loan, error)
	Return(id uuid.UUID) (domain.Loan, *domain.Fine, error)
	ListByMember(memberID uuid.UUID, status *domain.LoanStatus) ([]domain.Loan, error)
}
