package transaction

type service struct {
	transRepo Repository
}

func NewService(transRepo Repository) *service {
	return &service{
		transRepo: transRepo,
	}
}
