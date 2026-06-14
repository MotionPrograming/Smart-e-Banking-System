package transaction // ফোল্ডার: backend/service/transaction/

import (
	"errors"
	"smart-e-banking/backend/domain"
)

func (s *service) GetTransactionByID(id int64) (*domain.Transaction, error) {

	tx, err := s.transactionRepo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}

	if tx == nil {
		return nil, errors.New("transaction not found")
	}

	return tx, nil
}
