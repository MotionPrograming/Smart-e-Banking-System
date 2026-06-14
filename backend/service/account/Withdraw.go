package account

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func (s *service) Withdraw(accountID int64, amount decimal.Decimal, reference string) error {

	if amount.LessThanOrEqual(decimal.Zero) {
		return fmt.Errorf("invalid amount")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	acc, err := s.accountRepo.LockAccountForUpdate(tx, accountID)
	if err != nil {
		return err
	}

	if acc.Balance.LessThan(amount) {
		return fmt.Errorf("insufficient funds")
	}

	acc.Balance = acc.Balance.Sub(amount)

	if err := s.accountRepo.UpdateAccount(acc); err != nil {
		return err
	}

	return tx.Commit()
}
