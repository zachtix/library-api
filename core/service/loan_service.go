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

const maxActiveLoans = 3

func (s *loanServiceImpl) Borrow(memberID uuid.UUID, barcode string) (domain.Loan, error) {
	now := s.clock.Now()
	var created domain.Loan

	err := s.tx.Tx(func(r outport.TxRepos) error {
		member, err := r.Members.FindByID(memberID)
		if err != nil {
			return err
		}

		if member.Status != domain.MemberStatusActive {
			return domain.ErrMemberSuspended
		}

		hasUnpaid, err := r.Fines.HasUnpaidByMember(memberID)
		if err != nil {
			return err
		}
		if hasUnpaid {
			return domain.ErrHasUnpaidFines
		}

		active, err := r.Loans.ListActiveByMember(memberID)
		if err != nil {
			return err
		}
		for _, l := range active {
			if l.IsOverdue(now) {
				return domain.ErrHasOverdueLoans
			}
		}
		if len(active) >= maxActiveLoans {
			return domain.ErrLoanLimitReached
		}

		copie, err := r.Books.FindCopieByBarcode(barcode)
		if err != nil {
			return err
		}
		if copie.Status != domain.CopyStatusAvailable {
			return domain.ErrCopyNotAvailable
		}

		created, err = r.Loans.Save(domain.Loan{
			ID:         s.idGen.NewID(),
			MemberID:   memberID,
			CopyID:     copie.ID,
			BorrowedAt: now,
			DueAt:      domain.DueAtFor(now),
		})
		if err != nil {
			return err
		}
		return r.Books.UpdateCopieStatus(copie.ID, domain.CopyStatusBorrowed)
	})
	if err != nil {
		return domain.Loan{}, err
	}

	return created, nil
}

func (s *loanServiceImpl) Renew(id uuid.UUID) (domain.Loan, error)

func (s *loanServiceImpl) Return(id uuid.UUID) (domain.Loan, *domain.Fine, error)

func (s *loanServiceImpl) ListByMember(memberID uuid.UUID, status *domain.LoanStatus) ([]domain.Loan, error)
