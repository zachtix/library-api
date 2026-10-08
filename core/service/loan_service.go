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

func (s *loanServiceImpl) Renew(id uuid.UUID) (domain.Loan, error) {
	loan, err := s.repo.FindByID(id)
	if err != nil {
		return domain.Loan{}, err
	}
	if loan.ReturnedAt != nil {
		return domain.Loan{}, domain.ErrLoanAlreadyReturned
	}
	if loan.IsOverdue(s.clock.Now()) {
		return domain.Loan{}, domain.ErrLoanOverdue
	}
	if loan.RenewCount > 0 {
		return domain.Loan{}, domain.ErrRenewLimitReached
	}

	loan.DueAt = domain.RenewedDueAt(loan.DueAt)
	loan.RenewCount = 1

	return s.repo.Update(loan)
}

func (s *loanServiceImpl) Return(id uuid.UUID) (domain.Loan, *domain.Fine, error) {
	now := s.clock.Now()
	var returned domain.Loan
	var fine *domain.Fine
	err := s.tx.Tx(func(r outport.TxRepos) error {
		loan, err := r.Loans.FindByID(id)
		if err != nil {
			return err
		}
		if loan.ReturnedAt != nil {
			return domain.ErrLoanAlreadyReturned
		}

		loan.ReturnedAt = &now
		returned, err = r.Loans.Update(loan)
		if err != nil {
			return err
		}

		if err := r.Books.UpdateCopieStatus(loan.CopyID, domain.CopyStatusAvailable); err != nil {
			return err
		}

		if amount := domain.CalculateFine(loan.DueAt, now); amount > 0 {
			saved, err := r.Fines.Save(domain.Fine{
				ID:       s.idGen.NewID(),
				LoanID:   loan.ID,
				MemberID: loan.MemberID,
				Amount:   amount,
			})
			if err != nil {
				return err
			}
			fine = &saved
		}
		return nil
	})
	if err != nil {
		return domain.Loan{}, nil, err
	}
	return returned, fine, nil
}

func (s *loanServiceImpl) ListByMember(memberID uuid.UUID, status *domain.LoanStatus) ([]domain.Loan, error) {
	loans, err := s.repo.ListByMember(memberID)
	if err != nil {
		return nil, err
	}
	if status == nil {
		return loans, nil
	}

	now := s.clock.Now()
	filtered := make([]domain.Loan, 0, len(loans))
	for _, l := range loans {
		if l.Status(now) == *status {
			filtered = append(filtered, l)
		}
	}
	return filtered, nil
}
