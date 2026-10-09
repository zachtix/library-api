package domain_test

import (
	"library/core/domain"
	"testing"
	"time"
)

func TestDueAtFor(t *testing.T) {
	tests := []struct {
		name     string
		borrowed time.Time
		want     time.Time
	}{
		{"morning borrow", mustParse("2026-10-09T08:00:00+07:00"), mustParse("2026-10-23T23:59:59+07:00")},
		{"late night borrow counts same day", mustParse("2026-10-09T23:59:00+07:00"), mustParse("2026-10-23T23:59:59+07:00")},
		{"UTC input already next day in Bangkok", mustParse("2026-10-09T17:30:00Z"), mustParse("2026-10-24T23:59:59+07:00")},
		{"crosses year boundary", mustParse("2026-12-25T10:00:00+07:00"), mustParse("2027-01-08T23:59:59+07:00")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := domain.DueAtFor(tt.borrowed); !got.Equal(tt.want) {
				t.Errorf("DueAtFor(%v) = %v, want %v", tt.borrowed, got, tt.want)
			}
		})
	}
}

func TestRenewedDueAt(t *testing.T) {
	tests := []struct {
		name string
		due  time.Time
		want time.Time
	}{
		{"adds 7 days", mustParse("2026-10-14T23:59:59+07:00"), mustParse("2026-10-21T23:59:59+07:00")},
		{"crosses month boundary", mustParse("2026-10-28T23:59:59+07:00"), mustParse("2026-11-04T23:59:59+07:00")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := domain.RenewedDueAt(tt.due); !got.Equal(tt.want) {
				t.Errorf("RenewedDueAt(%v) = %v, want %v", tt.due, got, tt.want)
			}
		})
	}
}

func TestLoanStatus(t *testing.T) {
	due := mustParse("2026-10-23T23:59:59+07:00")
	returnedOnTime := mustParse("2026-10-20T10:00:00+07:00")
	returnedLate := mustParse("2026-10-30T10:00:00+07:00")

	tests := []struct {
		name       string
		returnedAt *time.Time
		now        time.Time
		want       domain.LoanStatus
	}{
		{"before due", nil, mustParse("2026-10-20T10:00:00+07:00"), domain.LoanStatusActive},
		{"last second of due day", nil, mustParse("2026-10-23T23:59:59+07:00"), domain.LoanStatusActive},
		{"first second of next day", nil, mustParse("2026-10-24T00:00:00+07:00"), domain.LoanStatusOverdue},
		{"UTC now already next day in Bangkok", nil, mustParse("2026-10-23T17:30:00Z"), domain.LoanStatusOverdue},
		{"returned on time, checked much later", &returnedOnTime, mustParse("2026-12-01T10:00:00+07:00"), domain.LoanStatusReturned},
		{"returned late still RETURNED", &returnedLate, mustParse("2026-12-01T10:00:00+07:00"), domain.LoanStatusReturned},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loan := domain.Loan{DueAt: due, ReturnedAt: tt.returnedAt}
			if got := loan.Status(tt.now); got != tt.want {
				t.Errorf("Status(%v) = %v, want %v", tt.now, got, tt.want)
			}
		})
	}
}

func TestLateDays(t *testing.T) {
	due := mustParse("2026-10-23T23:59:59+07:00")

	tests := []struct {
		name string
		at   time.Time
		want int
	}{
		{"3 days early is negative", mustParse("2026-10-20T10:00:00+07:00"), -3},
		{"same day", mustParse("2026-10-23T10:00:00+07:00"), 0},
		{"next day", mustParse("2026-10-24T10:00:00+07:00"), 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := domain.LateDays(due, tt.at); got != tt.want {
				t.Errorf("LateDays(%v, %v) = %d, want %d", due, tt.at, got, tt.want)
			}
		})
	}
}
