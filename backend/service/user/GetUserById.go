package user

import "smart-e-banking/backend/domain"

func (svc *service) GetUserByID(id int64) (*domain.User, error) {
	usr, err := svc.usrRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
