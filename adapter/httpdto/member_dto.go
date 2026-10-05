package httpdto

import (
	"library/core/domain"
	"time"

	"github.com/google/uuid"
)

type CreateMemberRequest struct {
	Name  string `json:"name" validate:"required,max=100"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateMemberStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=ACTIVE SUSPENDED"`
}

type MemberResponse struct {
	ID        uuid.UUID           `json:"id"`
	Name      string              `json:"name"`
	Email     string              `json:"email"`
	Status    domain.MemberStatus `json:"status"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

func (r CreateMemberRequest) MemberToDomain() domain.Member {
	return domain.Member{
		Name:  r.Name,
		Email: r.Email,
	}
}

func MemberResponseFromDomain(m domain.Member) MemberResponse {
	return MemberResponse{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
