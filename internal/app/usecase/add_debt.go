package usecase

import "context"

type AddDebtUseCase struct{}

func NewAddDebtUseCase() *AddDebtUseCase {
	return &AddDebtUseCase{}
}

func (uc *AddDebtUseCase) Execute(ctx context.Context) error {
	// Implement the logic to add debt here
	return nil
}
