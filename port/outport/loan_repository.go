package outport

import (
	"library/core/domain"

	"github.com/google/uuid"
)

type LoanRepository interface {
	Save(loan domain.Loan) (domain.Loan, error)
	FindByID(id uuid.UUID) (domain.Loan, error)
	Update(loan domain.Loan) (domain.Loan, error)
	ListActiveByMember(memberID uuid.UUID) ([]domain.Loan, error) // return not yet
	ListByMember(memberID uuid.UUID) ([]domain.Loan, error)
}
