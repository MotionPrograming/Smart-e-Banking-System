package account

import "errors"

func (s *service) DeleteAccount(id int64) error {

	// 1. check account exists
	acc, err := s.accountRepo.GetAccountByID(id)
	if err != nil {
		return err
	}

	// 2. business rule: balance check
	if !acc.Balance.IsZero() {
		return errors.New("cannot delete account with balance")
	}

	// 3. call repo (soft delete recommended)
	return s.accountRepo.DeleteAccount(id)
}
