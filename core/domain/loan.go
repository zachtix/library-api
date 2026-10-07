package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	LoanPeriodDays  = 14
	RenewPeriodDays = 7
)

var LibraryLocation = time.FixedZone("Asia/Bangkok", 7*60*60)

type Loan struct {
	ID         uuid.UUID
	MemberID   uuid.UUID
	CopyID     uuid.UUID
	BorrowedAt time.Time
	DueAt      time.Time
	ReturnedAt *time.Time
	RenewCount int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

var (
	ErrMemberSuspended     = errors.New("member is suspended")
	ErrHasUnpaidFines      = errors.New("member has unpaid fines")
	ErrHasOverdueLoans     = errors.New("member has overdue loans")
	ErrLoanLimitReached    = errors.New("member already has 3 active loans")
	ErrCopyNotFound        = errors.New("copy not found")
	ErrCopyNotAvailable    = errors.New("copy is not available")
	ErrLoanNotFound        = errors.New("loan not found")
	ErrLoanAlreadyReturned = errors.New("loan already returned")
	ErrLoanOverdue         = errors.New("loan is overdue")
	ErrRenewLimitReached   = errors.New("loan has already been renewed")
)

type LoanStatus string

const (
	LoanStatusActive   LoanStatus = "ACTIVE"
	LoanStatusOverdue  LoanStatus = "OVERDUE"
	LoanStatusReturned LoanStatus = "RETURNED"
)

func (l Loan) Status(now time.Time) LoanStatus {
	if l.ReturnedAt != nil {
		return LoanStatusReturned
	}
	if l.IsOverdue(now) {
		return LoanStatusOverdue
	}
	return LoanStatusActive
}

func startOfDay(t time.Time) time.Time {
	y, m, d := t.In(LibraryLocation).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, LibraryLocation)
}

func endOfDay(day time.Time) time.Time {
	return day.Add(24*time.Hour - time.Second)
}

func LateDays(dueAt, t time.Time) int {
	return int(startOfDay(t).Sub(startOfDay(dueAt)).Hours() / 24)
}

func (l Loan) IsOverdue(now time.Time) bool {
	return LateDays(l.DueAt, now) > 0
}

func DueAtFor(borrowedAt time.Time) time.Time {
	return endOfDay(startOfDay(borrowedAt).AddDate(0, 0, LoanPeriodDays))
}

func RenewedDueAt(dueAt time.Time) time.Time {
	return endOfDay(startOfDay(dueAt).AddDate(0, 0, RenewPeriodDays))
}
