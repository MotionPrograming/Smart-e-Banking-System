package account

import "errors"

func (r *accountRepository) DeleteAccount(id int64) error {

	query := `
		UPDATE accounts
		SET status = 'closed'
		WHERE id = ? AND status = 'active'
	`

	result, err := r.db.Exec(query, id)
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
