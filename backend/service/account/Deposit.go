package account

import (
	"errors"

	"github.com/shopspring/decimal"
)

func (s *service) Deposit(accountID int64, amount decimal.Decimal) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("deposit amount must be greater than zero")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	acc, err := s.accountRepo.LockAccountForUpdate(tx, accountID)
	if err != nil {
		return err
	}
	if acc.Status != "active" {
		err = errors.New("account is not active")
		return err
	}

	acc.Balance = acc.Balance.Add(amount)

	if err = s.accountRepo.UpdateBalance(tx, acc); err != nil {
		return err
	}

	// Record the transaction inside the same tx so it rolls back on failure.
	_, err = tx.Exec(`
		INSERT INTO transactions
		    (from_account_id, to_account_id, amount, currency, type, status)
		VALUES (NULL, ?, ?, ?, 'deposit', 'completed')
	`, accountID, amount, acc.Currency)
	if err != nil {
		return err
	}

	return tx.Commit()
}
