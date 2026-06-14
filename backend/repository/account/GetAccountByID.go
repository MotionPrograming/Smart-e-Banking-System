package account

import (
	"database/sql"
	"errors"
	"smart-e-banking/backend/domain"
)

func (r *accountRepository) GetAccountByID(accountID int64) (*domain.Account, error) {
	var acc domain.Account
	err := r.db.Get(&acc, `
		SELECT id, user_id, account_number, balance, currency,
		       account_type, status, created_at, updated_at
		FROM accounts
		WHERE id = ? AND status = 'active'
	`, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("account not found or inactive")
		}
		return nil, err
	}
	return &acc, nil
}
