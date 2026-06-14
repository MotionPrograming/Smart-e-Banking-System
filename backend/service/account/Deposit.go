package account

import (
	"errors"

	"github.com/shopspring/decimal"
)

func (s *service) Deposit(accountID int64, amount decimal.Decimal) error {

	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("invalid deposit amount")
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

	// 🔒 lock account (important for banking safety)
	acc, err := s.accountRepo.LockAccountForUpdate(tx, accountID)
	if err != nil {
		return err
	}

	if acc.Status != "active" {
		return errors.New("account is not active")
	}

	acc.Balance = acc.Balance.Add(amount)

	err = s.accountRepo.UpdateAccount(acc)
	if err != nil {
		return err
	}

	return tx.Commit()
}
