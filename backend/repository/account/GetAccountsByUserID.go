package account

import "smart-e-banking/backend/domain"

func (r *accountRepository) GetAccountsByUserID(userID int64) ([]domain.Account, error) {
	var accounts []domain.Account

	query := `
		SELECT
			id,
			user_id,
			account_number,
			balance,
			currency,
			account_type,
			status,
			created_at,
			updated_at
		FROM accounts
		WHERE user_id = ? AND status = 'active'
		ORDER BY id DESC
	`

	err := r.db.Select(&accounts, query, userID)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
