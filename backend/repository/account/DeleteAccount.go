package account

import "errors"

func (r *accountRepository) DeleteAccount(id int64) error {
	result, err := r.db.Exec(
		`UPDATE accounts SET status = 'closed', updated_at = NOW()
		 WHERE id = ? AND status = 'active'`,
		id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("account not found or already closed")
	}
	return nil
}
