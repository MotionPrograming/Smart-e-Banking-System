package transaction

import (
	"database/sql"
	"errors"
	"smart-e-banking/backend/domain"
)

func (s *service) GetTransactionByID(id int64) (*domain.Transaction, error) {
	if id <= 0 {
		return nil, errors.New("invalid transaction id")
	}
	tx, err := s.transactionRepo.GetTransactionByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("transaction not found")
		}
		return nil, err
	}
	return tx, nil
}
