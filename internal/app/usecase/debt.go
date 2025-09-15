package usecase

import (
	"line-bot/internal/platform/database"
)

type DebtUseCase struct {
	repo database.DebtPostgres
}

func NewDebtUseCase(repo database.DebtPostgres) *DebtUseCase {
	return &DebtUseCase{repo: repo}
}

func (uc *DebtUseCase) GetDebts() (string, error) {
	debt, err := uc.repo.GetDebts()
	if err != nil {
		return "", err
	}
	return debt, nil
}
