package user

import (
	"errors"
	"smart-e-banking/backend/domain"

	"golang.org/x/crypto/bcrypt"
)

func (svc *service) Authenticate(email, password string) (*domain.User, error) {

	usr, err := svc.usrRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(usr.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return usr, nil
}
