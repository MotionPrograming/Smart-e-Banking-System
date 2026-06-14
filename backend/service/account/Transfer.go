package account

import (
	"errors"
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

func (s *service) Transfer(fromAccountID, toAccountID int64, amount decimal.Decimal) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("transfer amount must be greater than zero")
	}
	if fromAccountID == toAccountID {
		return errors.New("cannot transfer to the same account")
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

	// Always lock in ascending ID order to prevent deadlocks.
	var fromAcc, toAcc *domain.Account
	if fromAccountID < toAccountID {
		fromAcc, err = s.accountRepo.LockAccountForUpdate(tx, fromAccountID)
		if err != nil {
			return err
		}
		toAcc, err = s.accountRepo.LockAccountForUpdate(tx, toAccountID)
		if err != nil {
			return err
		}
	} else {
		toAcc, err = s.accountRepo.LockAccountForUpdate(tx, toAccountID)
		if err != nil {
			return err
		}
		fromAcc, err = s.accountRepo.LockAccountForUpdate(tx, fromAccountID)
		if err != nil {
			return err
		}
	}

	if fromAcc.Status != "active" || toAcc.Status != "active" {
		err = errors.New("one or both accounts are not active")
		return err
	}
	if fromAcc.Balance.LessThan(amount) {
		err = errors.New("insufficient balance")
		return err
	}

	fromAcc.Balance = fromAcc.Balance.Sub(amount)
	toAcc.Balance = toAcc.Balance.Add(amount)

	if err = s.accountRepo.UpdateBalance(tx, fromAcc); err != nil {
		return err
	}
	if err = s.accountRepo.UpdateBalance(tx, toAcc); err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO transactions
		    (from_account_id, to_account_id, amount, currency, type, status)
		VALUES (?, ?, ?, ?, 'transfer', 'completed')
	`, fromAccountID, toAccountID, amount, fromAcc.Currency)
	if err != nil {
		return err
	}

	return tx.Commit()
}
