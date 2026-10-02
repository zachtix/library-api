package gormadapter

import (
	"library/core/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberModel struct {
	ID        uuid.UUID           `gorm:"primaryKey;type:uuid;not null"`
	Name      string              `gorm:"not null"`
	Email     string              `gorm:"uniqueIndex:idx_members_email,where:deleted_at IS NULL;not null"`
	Status    domain.MemberStatus `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (MemberModel) TableName() string { return "members" }

func (m MemberModel) toDomain() domain.Member {
	return domain.Member{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: deletedAtToDomain(m.DeletedAt),
	}
}

func memberFromDomain(m domain.Member) MemberModel {
	return MemberModel{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: deletedAtFromDomain(m.DeletedAt),
	}
}
