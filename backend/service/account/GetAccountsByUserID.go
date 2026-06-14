package account

import "smart-e-banking/backend/domain"

func (s *service) GetAccountsByUserID(userID int64) ([]domain.Account, error) {

	accounts, err := s.accountRepo.GetAccountsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var activeAccounts []domain.Account

	for _, acc := range accounts {
		if acc.Status == "active" {
			activeAccounts = append(activeAccounts, acc)
		}
	}

	return activeAccounts, nil
}
