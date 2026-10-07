package gormadapter

import (
	"library/port/outport"

	"gorm.io/gorm"
)

type GormTx struct {
	db *gorm.DB
}

func NewGormTx(db *gorm.DB) outport.TxManager {
	return &GormTx{db: db}
}

func (m *GormTx) Tx(fn func(r outport.TxRepos) error) error {
	return m.db.Transaction(func(gtx *gorm.DB) error {
		return fn(outport.TxRepos{
			Members: NewGormMemberRepository(gtx),
			Books:   NewGormBookRepository(gtx),
			Loans:   NewGormLoanRepository(gtx),
			Fines:   NewGormFineRepository(gtx),
		})
	})
}
