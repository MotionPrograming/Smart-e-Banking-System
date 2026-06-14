package transaction

import (
	"smart-e-banking/backend/domain"
)

func (r *transactionRepository) CreateTransaction(tx domain.Transaction) (*domain.Transaction, error) {
	if tx.Status == "" {
		tx.Status = "completed"
	}
	if tx.Currency == "" {
		tx.Currency = "BDT"
	}

	result, err := r.db.Exec(`
		INSERT INTO transactions
		    (from_account_id, to_account_id, amount, currency, type, status, reference)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`,
		tx.FromAccountID,
		tx.ToAccountID,
		tx.Amount,
		tx.Currency,
		tx.Type,
		tx.Status,
		tx.Reference,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	tx.ID = id
	return &tx, nil
}
