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

func (r *GormFineRepository) Save(fine domain.Fine) (domain.Fine, error)
func (r *GormFineRepository) HasUnpaidByMember(memberID uuid.UUID) (bool, error)
