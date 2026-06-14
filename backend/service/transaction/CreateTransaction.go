package transaction

import (
	"errors"
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

func (s *service) CreateTransaction(tx domain.Transaction) (*domain.Transaction, error) {
	if tx.Amount.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("transaction amount must be greater than zero")
	}
	if tx.Type == "" {
		return nil, errors.New("transaction type is required")
	}
	if tx.Status == "" {
		tx.Status = "completed"
	}
	return s.transactionRepo.CreateTransaction(tx)
}
