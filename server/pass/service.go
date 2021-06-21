package pass

type Service interface {
}

type service struct {
	repository Repository
}

func NewPassService(repository Repository) *service {
	return &service{repository}
}
