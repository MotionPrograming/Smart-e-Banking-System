package account

type Service struct {
	repository *repository
}

func NewService(repository *repository) *Service {
	return &Service{
		repository: repository,
	}
}

type repository struct {
	accRepo accountRepository
}
