package domain

import (
	"errors"
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

var (
	ErrMemberNotFound = errors.New("member not found")
	ErrEmailTaken     = errors.New("email already taken")
	ErrInvalidStatus  = errors.New("invalid member status")
)

func (s MemberStatus) Valid() bool {
	switch s {
	case MemberStatusActive, MemberStatusSuspended:
		return true
	}
	return false
}
