package transaction // ফোল্ডার: backend/service/transaction/

import (
	"errors"
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

func (h *service) CreateTransaction(tx domain.Transaction) (*domain.Transaction, error) {

	// 🔐 validation
	if tx.Amount.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("invalid transaction amount")
	}

	if tx.Type == "" {
		return nil, errors.New("transaction type required")
	}

	if tx.Status == "" {
		tx.Status = "completed"
	}

	return h.transactionRepo.CreateTransaction(tx)
}
