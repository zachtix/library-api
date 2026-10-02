package domain

import (
	"time"

	"github.com/google/uuid"
)

type MemberStatus string

const (
	MemberStatusActive    MemberStatus = "ACTIVE"
	MemberStatusSuspended MemberStatus = "SUSPENDED"
)

type Member struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Status    MemberStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
