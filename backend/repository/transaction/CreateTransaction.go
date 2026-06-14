package transaction

import (
	"smart-e-banking/backend/domain"
)

func (r *transactionRepository) CreateTransaction(tx domain.Transaction) (*domain.Transaction, error) {

	query := `
		INSERT INTO transactions 
		(from_account_id, to_account_id, amount, currency, type, status, reference)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.Exec(
		query,
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
