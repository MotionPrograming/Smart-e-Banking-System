package user

import "smart-e-banking/backend/domain"

func (svc *service) CreateUser(user domain.User) (*domain.User, error) {
	usr, err := svc.usrRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
