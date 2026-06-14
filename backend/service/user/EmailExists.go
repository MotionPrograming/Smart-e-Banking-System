package user

func (svc *service) EmailExists(email string) (bool, error) {
	exists, err := svc.usrRepo.EmailExists(email)
	if err != nil {
		return false, err
	}
	return exists, nil
}
