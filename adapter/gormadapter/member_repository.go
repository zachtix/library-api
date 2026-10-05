package gormadapter

import (
	"errors"
	"library/core/domain"
	"library/port/outport"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormMemberRepository struct {
	db *gorm.DB
}

func NewGormMemberRepository(db *gorm.DB) outport.MemberRepository {
	return &GormMemberRepository{
		db: db,
	}
}

func (r *GormMemberRepository) Save(member domain.Member) (domain.Member, error) {
	model := memberFromDomain(member)
	if err := r.db.Create(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.Member{}, domain.ErrEmailTaken
		}
		return domain.Member{}, err
	}
	return model.toDomain(), nil
}
func (r *GormMemberRepository) FindByID(id uuid.UUID) (domain.Member, error) {
	var model MemberModel
	if err := r.db.Take(&model, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Member{}, domain.ErrMemberNotFound
		}
		return domain.Member{}, err
	}
	return model.toDomain(), nil
}
func (r *GormMemberRepository) SetStatus(id uuid.UUID, status domain.MemberStatus) (domain.Member, error) {
	res := r.db.Model(&MemberModel{}).Where("id = ?", id).Update("status", status)
	if res.Error != nil {
		return domain.Member{}, res.Error
	}
	if res.RowsAffected == 0 {
		return domain.Member{}, domain.ErrMemberNotFound
	}
	return r.FindByID(id)
}
