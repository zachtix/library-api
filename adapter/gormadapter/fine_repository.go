package gormadapter

import (
	"library/core/domain"
	"library/port/outport"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormFineRepository struct {
	db *gorm.DB
}

func NewGormFineRepository(db *gorm.DB) outport.FineRepository {
	return &GormFineRepository{
		db: db,
	}
}

func (r *GormFineRepository) Save(fine domain.Fine) (domain.Fine, error) {
	model := fineFromDomain(fine)
	if err := r.db.Create(&model).Error; err != nil {
		return domain.Fine{}, err
	}
	return model.toDomain(), nil
}
func (r *GormFineRepository) HasUnpaidByMember(memberID uuid.UUID) (bool, error) {
	var count int64
	if err := r.db.Model(&FineModel{}).Where("member_id = ? AND paid_at IS NULL", memberID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
