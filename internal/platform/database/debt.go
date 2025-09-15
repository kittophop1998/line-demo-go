package database

type DebtPostgres struct{}

func NewDebtRepo() DebtPostgres {
	return DebtPostgres{}
}

func (r *DebtPostgres) GetDebts() (string, error) {
	return "1000 THB", nil
}
