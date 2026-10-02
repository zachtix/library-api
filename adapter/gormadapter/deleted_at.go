package gormadapter

import (
	"time"

	"gorm.io/gorm"
)

func deletedAtToDomain(d gorm.DeletedAt) *time.Time {
	if !d.Valid {
		return nil
	}
	t := d.Time
	return &t
}

func deletedAtFromDomain(t *time.Time) gorm.DeletedAt {
	if t == nil {
		return gorm.DeletedAt{}
	}
	return gorm.DeletedAt{Time: *t, Valid: true}
}
