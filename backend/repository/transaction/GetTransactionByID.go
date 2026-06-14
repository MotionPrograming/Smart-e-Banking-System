package transaction

import (
	"smart-e-banking/backend/domain"
)

func (r *transactionRepository) GetTransactionByID(transactionID int64) (*domain.Transaction, error) {
	var tx domain.Transaction
	err := r.db.Get(&tx, `
		SELECT id, from_account_id, to_account_id, amount, currency,
		       type, status, reference, created_at
		FROM transactions
		WHERE id = ?
	`, transactionID)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}
