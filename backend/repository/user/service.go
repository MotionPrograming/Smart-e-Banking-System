package user

type service struct {
	userRepo Repository
}

func NewService(userRepo Repository) *service {
	return &service{
		userRepo: userRepo,
	}
}
