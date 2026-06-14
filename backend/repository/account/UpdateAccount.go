package account

import (
	"smart-e-banking/backend/domain"

	"github.com/jmoiron/sqlx"
)

// UpdateBalance writes the balance inside a caller-managed transaction.
// Always use this from Deposit / Withdraw / Transfer.
func (r *accountRepository) UpdateBalance(tx *sqlx.Tx, acc *domain.Account) error {
	_, err := tx.Exec(
		`UPDATE accounts SET balance = ?, updated_at = NOW() WHERE id = ?`,
		acc.Balance, acc.ID,
	)
	return err
}

// UpdateAccount updates non-balance fields (account_type, status).
func (r *accountRepository) UpdateAccount(acc *domain.Account) error {
	_, err := r.db.Exec(
		`UPDATE accounts SET account_type = ?, status = ?, updated_at = NOW() WHERE id = ?`,
		acc.AccountType, acc.Status, acc.ID,
	)
	return err
}
