package service

import (
	"library/core/domain"
	"library/port/inport"
	"library/port/outport"

	"github.com/google/uuid"
)

type loanServiceImpl struct {
	repo  outport.LoanRepository
	idGen outport.IDGenerator
	tx    outport.TxManager
	clock outport.Clock
}

func NewLoanService(repo outport.LoanRepository, idGen outport.IDGenerator, tx outport.TxManager, clock outport.Clock) inport.LoanService {
	return &loanServiceImpl{
		repo:  repo,
		idGen: idGen,
		tx:    tx,
		clock: clock,
	}
}

func (s *loanServiceImpl) Borrow(memberID uuid.UUID, barcode string) (domain.Loan, error)

func (s *loanServiceImpl) Renew(id uuid.UUID) (domain.Loan, error)

func (s *loanServiceImpl) Return(id uuid.UUID) (domain.Loan, *domain.Fine, error)

func (s *loanServiceImpl) ListByMember(memberID uuid.UUID, status *domain.LoanStatus) ([]domain.Loan, error)
