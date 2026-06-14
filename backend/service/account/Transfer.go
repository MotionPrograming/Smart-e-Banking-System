package account

import (
	"fmt"
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

func (s *service) Transfer(fromAccountID, toAccountID int64, amount decimal.Decimal) error {

	if amount.LessThanOrEqual(decimal.Zero) {
		return fmt.Errorf("invalid amount")
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

	var fromAccount, toAccount *domain.Account

	// 🔒 deadlock-safe lock order
	if fromAccountID < toAccountID {

		fromAccount, err = s.accountRepo.LockAccountForUpdate(tx, fromAccountID)
		if err != nil {
			return err
		}

		toAccount, err = s.accountRepo.LockAccountForUpdate(tx, toAccountID)
		if err != nil {
			return err
		}

	} else {

		toAccount, err = s.accountRepo.LockAccountForUpdate(tx, toAccountID)
		if err != nil {
			return err
		}

		fromAccount, err = s.accountRepo.LockAccountForUpdate(tx, fromAccountID)
		if err != nil {
			return err
		}
	}

	if fromAccount == nil || toAccount == nil {
		return fmt.Errorf("account not found")
	}

	if fromAccount.Balance.LessThan(amount) {
		return fmt.Errorf("insufficient balance")
	}

	fromAccount.Balance = fromAccount.Balance.Sub(amount)
	toAccount.Balance = toAccount.Balance.Add(amount)

	err = s.accountRepo.UpdateAccount(fromAccount)
	if err != nil {
		return err
	}

	err = s.accountRepo.UpdateAccount(toAccount)
	if err != nil {
		return err
	}

	return tx.Commit()
}
