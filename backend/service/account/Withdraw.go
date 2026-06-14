package account

import (
	"errors"

	"github.com/shopspring/decimal"
)

func (s *service) Withdraw(accountID int64, amount decimal.Decimal, reference string) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("withdrawal amount must be greater than zero")
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
	if acc.Balance.LessThan(amount) {
		err = errors.New("insufficient funds")
		return err
	}

	acc.Balance = acc.Balance.Sub(amount)

	if err = s.accountRepo.UpdateBalance(tx, acc); err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO transactions
		    (from_account_id, to_account_id, amount, currency, type, status, reference)
		VALUES (?, NULL, ?, ?, 'withdraw', 'completed', ?)
	`, accountID, amount, acc.Currency, reference)
	if err != nil {
		return err
	}

	return tx.Commit()
}
