package domain_test

import (
	"library/core/domain"
	"testing"
	"time"
)

// Helper RFC 3339
func mustParse(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

func TestCalculateFine(t *testing.T) {
	due := mustParse("2026-10-23T23:59:59+07:00")

	tests := []struct {
		name     string
		returned time.Time
		want     int64
	}{
		{"returned early", mustParse("2026-10-20T10:00:00+07:00"), 0},
		{"last second of due day", mustParse("2026-10-23T23:59:59+07:00"), 0},
		{"first second of next day", mustParse("2026-10-24T00:00:00+07:00"), 10},
		{"3 days late", mustParse("2026-10-26T15:00:00+07:00"), 30},
		{"49 days late, just under cap", mustParse("2026-12-11T09:00:00+07:00"), 490},
		{"50 days late, hits cap", mustParse("2026-12-12T09:00:00+07:00"), 500},
		{"far beyond cap", mustParse("2027-03-01T09:00:00+07:00"), 500},
		{"UTC input already next day in Bangkok", mustParse("2026-10-23T17:30:00Z"), 10},
		{"last second of first late day", mustParse("2026-10-24T23:59:59+07:00"), 10},
		{"same day before due time", mustParse("2026-10-23T08:00:00+07:00"), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := domain.CalculateFine(due, tt.returned); got != tt.want {
				t.Errorf("CalculateFine(%v, %v) = %d, want %d", due, tt.returned, got, tt.want)
			}
		})
	}
}
