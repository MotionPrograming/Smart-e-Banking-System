package user

import "smart-e-banking/backend/domain"

func (svc *service) GetUserByEmail(email string) (*domain.User, error) {
	usr, err := svc.usrRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
