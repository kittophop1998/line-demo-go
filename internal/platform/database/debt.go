package database

import "gorm.io/gorm"

type DebtPostgres struct {
	db *gorm.DB
}

func NewDebtRepo(db *gorm.DB) DebtPostgres {
	return DebtPostgres{db: db}
}

func (r *DebtPostgres) GetDebts() (string, error) {
	return "1000 THB", nil
}
