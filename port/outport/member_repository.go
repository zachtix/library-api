package outport

import (
	"library/core/domain"

	"github.com/google/uuid"
)

type MemberRepository interface {
	Save(member domain.Member) (domain.Member, error)
	FindByID(id uuid.UUID) (domain.Member, error)
	SetStatus(id uuid.UUID, status domain.MemberStatus) (domain.Member, error)
}
