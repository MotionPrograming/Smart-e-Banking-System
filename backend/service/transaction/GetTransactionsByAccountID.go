package transaction

import (
	"errors"
	"smart-e-banking/backend/domain"
)

func (s *service) GetTransactionsByAccountID(accountID int64) ([]domain.Transaction, error) {
	if accountID <= 0 {
		return nil, errors.New("invalid account id")
	}
	return s.transactionRepo.GetTransactionsByAccountID(accountID)
}
