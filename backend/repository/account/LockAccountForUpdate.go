package account

import (
	"database/sql"
	"errors"
	"smart-e-banking/backend/domain"

	"github.com/jmoiron/sqlx"
)

func (r *accountRepository) LockAccountForUpdate(tx *sqlx.Tx, id int64) (*domain.Account, error) {
	var acc domain.Account
	err := tx.Get(&acc, `SELECT * FROM accounts WHERE id = ? FOR UPDATE`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("account not found")
		}
		return nil, err
	}
	return &acc, nil
}
