package repository

type DebtPostgres interface {
	GetDebts() (string, error)
}
