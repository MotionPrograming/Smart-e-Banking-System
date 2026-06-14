package transaction

import (
	"smart-e-banking/backend/domain"
)

func (r *transactionRepository) GetTransactionsByAccountID(accountID int64) ([]domain.Transaction, error) {

	query := `
		SELECT id, from_account_id, to_account_id, amount, currency, type, status, reference, created_at
		FROM transactions
		WHERE from_account_id = ? OR to_account_id = ?
		ORDER BY created_at DESC
	`

	var transactions []domain.Transaction
	err := r.db.Select(&transactions, query, accountID, accountID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
