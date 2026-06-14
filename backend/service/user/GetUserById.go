package user

import (
	"errors"
	"smart-e-banking/backend/domain"
)

func (svc *service) GetUserByID(id int64) (*domain.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user id")
	}
	usr, err := svc.usrRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
