package account

import (
	"smart-e-banking/backend/domain"
)

func (r *accountRepository) UpdateAccount(acc *domain.Account) error {
	query := `UPDATE accounts SET balance=$1 WHERE id=$2`
	_, err := r.db.Exec(query, acc.Balance, acc.ID)
	return err
}
