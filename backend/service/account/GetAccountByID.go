package account

import (
	"errors"
	"smart-e-banking/backend/domain"
)

func (s *service) GetAccountByID(id int64) (*domain.Account, error) {

	acc, err := s.accountRepo.GetAccountByID(id)
	if err != nil {
		return nil, err
	}

	if acc.Status != "active" {
		return nil, errors.New("account is not active")
	}

	return acc, nil
}
