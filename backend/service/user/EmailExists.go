package user

func (svc *service) EmailExists(email string) (bool, error) {
	return svc.usrRepo.EmailExists(email)
}
