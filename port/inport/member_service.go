package inport

import (
	"library/core/domain"

	"github.com/google/uuid"
)

type MemberService interface {
	Create(member domain.Member) (domain.Member, error)
	Get(id uuid.UUID) (domain.Member, error)
	UpdateStatus(id uuid.UUID, status domain.MemberStatus) (domain.Member, error)
}
